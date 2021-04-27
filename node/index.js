const git = require("nodegit");
const execa = require("execa");
const fs = require("fs");
const yaml = require("js-yaml");

const CORE_UI = "core-ui";
const CORE = "core";
const CATALOG = "catalog";
const BACKEND = "backend";

const BUSOLA_PATH = "/Users/i515358/Sites/kyma-project/busola";

const fileChangeMappings = [
  {
    prefix: "core-ui",
    bump: [CORE_UI],
  },
  {
    prefix: "core",
    bump: [CORE],
  },
  {
    prefix: "service-catalog-ui",
    bump: [CATALOG],
  },
  {
    prefix: "backend",
    bump: [BACKEND],
  },
  {
    prefix: "shared",
    bump: [CORE_UI, CATALOG],
  },
];

const bumpMappings = {
  [CORE_UI]: {
    filePath: BUSOLA_PATH + "/resources/web/deployment.yaml",
    registryPath: "eu.gcr.io/kyma-project/busola-backend:PR-",
    update: (content, image) => {
      content.spec.template.spec.containers.find(
        (c) => c.name === "core-ui"
      ).image = image;
    },
  },
  [CORE]: {
    filePath: BUSOLA_PATH + "/resources/web/deployment.yaml",
    registryPath: "eu.gcr.io/kyma-project/busola-core:PR-",
    update: (content, image) => {
      content.spec.template.spec.containers.find(
        (c) => c.name === "busola"
      ).image = image;
    },
  },
  [CATALOG]: {
    filePath: BUSOLA_PATH + "/resources/web/deployment-service-catalog.yaml",
    registryPath: "eu.gcr.io/kyma-project/busola-service-catalog-ui:PR-",
    update: (content, image) => {
      content.spec.template.spec.containers[0].image = image;
    },
  },
  [BACKEND]: {
    filePath: BUSOLA_PATH + "/resources/backend/deployment.yaml",
    registryPath: "eu.gcr.io/kyma-project/busola-backend:PR-",
    update: (content, image) => {
      content.spec.template.spec.containers[0].image = image;
    },
  },
};

const getPRNumber = async () => {
  try {
    const { stdout } = await execa("gh", `pr status --json url`.split(" "), {
      cwd: BUSOLA_PATH,
    });
    const pr = JSON.parse(stdout).currentBranch;
    if (!pr) {
      console.log("Looks like there's no PR for this branch. Sorry.");
      process.exit(1);
    } else {
      return pr.url.substr(pr.url.lastIndexOf("/") + 1);
    }
  } catch (e) {
    console.log("Whoops, cannot get PR number. Is gh installed?");
    console.warn(e.message);
    process.exit(1);
  }
};

(async () => {
  const prNumber = process.argv.slice(2)[0] || (await getPRNumber());

  const repo = await git.Repository.open(BUSOLA_PATH);
  const status = await repo.getStatus();
  const filePaths = status.map((s) => s.path());

  const toBump = new Set();
  for (const filePath of filePaths) {
    for (const mapping of fileChangeMappings) {
      if (filePath.startsWith(mapping.prefix)) {
        for (const b of mapping.bump) {
          toBump.add(b);
        }
      }
    }
  }

  if (toBump.size === 0) {
      console.log('nothing to bump')
      return;
  }

  for (const bump of [...toBump.keys()]) {
    console.log('bumping ' + bump)
    const mapping = bumpMappings[bump];
    const fileContent = fs.readFileSync(mapping.filePath);
    const yamlContent = yaml.load(fileContent);
    mapping.update(yamlContent, mapping.registryPath + prNumber);
    fs.writeFileSync(mapping.filePath, yaml.dump(yamlContent));
  }
})();
