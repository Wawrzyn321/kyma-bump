package mappings

func GetMappingPresets() Mappings {
	return Mappings{
		Mapping{
			Name:        "addons",
			Aliases:     []string{"addons", "add-ons", "a"},
			FilePath:    "resources/helm-broker/charts/addons-ui/values.yaml",
			YamlPath:    "image.version",
			RegistryUrl: "eu.gcr.io/kyma-project/add-ons-ui",
			Dir:         "add-ons/",
		},
		Mapping{
			Name:        "core",
			Aliases:     []string{"core", "ng"},
			FilePath:    "resources/console/charts/web/values.yaml",
			YamlPath:    "console.image.tag",
			RegistryUrl: "eu.gcr.io/kyma-project/console",
			Dir:         "core/",
		},
		Mapping{
			Name:        "core-ui",
			Aliases:     []string{"core-ui", "c"},
			FilePath:    "resources/console/charts/web/values.yaml",
			YamlPath:    "core_ui.image.tag",
			RegistryUrl: "eu.gcr.io/kyma-project/core-ui",
			Dir:         "core-ui/",
		},
		Mapping{
			Name:        "content",
			Aliases:     []string{"content", "docs", "d"},
			FilePath:    "resources/core/charts/docs/charts/content-ui/Chart.yaml",
			YamlPath:    "appVersion",
			RegistryUrl: "eu.gcr.io/kyma-project/content-ui",
			Dir:         "content/",
		},
		Mapping{
			Name:        "logging",
			Aliases:     []string{"logging", "logs", "l"},
			FilePath:    "resources/logging/charts/logui/values.yaml",
			YamlPath:    "image.tag",
			RegistryUrl: "eu.gcr.io/kyma-project/log-ui",

			Dir: "logging/",
		},
		Mapping{
			Name:        "service-catalog-ui",
			Aliases:     []string{"service-catalog-ui", "catalog", "sc"},
			FilePath:    "resources/service-catalog-addons/charts/service-catalog-ui/values.yaml",
			YamlPath:    "image.tag",
			RegistryUrl: "eu.gcr.io/kyma-project/service-catalog-ui",
			Dir:         "service-catalog-ui/",
		},
		Mapping{
			Name:        "tests",
			Aliases:     []string{"tests", "ui-test"},
			FilePath:    "resources/console/values.yaml",
			YamlPath:    "global.ui_acceptance_tests.version",
			RegistryUrl: "eu.gcr.io/kyma-project/ui-tests",
			Dir:         "tests/",
		},
		Mapping{
			Name:        "console-backend-service",
			Aliases:     []string{"console-backend-service", "cbs"},
			FilePath:    "resources/console/values.yaml",
			YamlPath:    "global.console_backend_service.version",
			RegistryUrl: "eu.gcr.io/kyma-project/console-backend-service",
			Dir:         "components/console-backend-service/",
		},
		Mapping{
			Name:        "console-backend-service-test",
			Aliases:     []string{"console-backend-service-test", "cbs-test"},
			FilePath:    "resources/console/values.yaml",
			YamlPath:    "global.console_backend_service_test.version",
			RegistryUrl: "eu.gcr.io/kyma-project/console-backend-service-test",
			Dir:         "tests/console-backend-service",
		},
	}
}
