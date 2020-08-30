package sysCommands

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func GetChangedFiles(revision string) ([]string, error) {
	repoPath := path.Join(os.Getenv("BUMP_KYMA_HOME"))
	cmdStr := fmt.Sprintf("git --git-dir=%s diff --name-only --no-index", repoPath)
	fmt.Println(cmdStr)
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	return strings.Split(string(out), "\n"), err
}
