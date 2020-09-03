package sysCommands

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetChangedFiles(repoPath string, revision *string) ([]string, error) {
	var cmdStr string
	if revision == nil {
		cmdStr = fmt.Sprintf("git --git-dir=%s/.git --work-tree=%s diff --name-only --no-index", repoPath, repoPath)
	} else {
		cmdStr = fmt.Sprintf("git --git-dir=%s/.git --work-tree=%s diff %s --name-only", repoPath, repoPath, *revision)
	}
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	return strings.Split(string(out), "\n"), err
}
