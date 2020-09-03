package sysCommands

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetChangedFiles(repoPath string, revision *string) ([]string, error) {
	var cmdStr string
	if revision == nil {
		cmdStr = fmt.Sprintf("git --git-dir=%s/.git --work-tree=%s diff --name-only", repoPath, repoPath)
	} else {
		cmdStr = fmt.Sprintf("git --git-dir=%s/.git --work-tree=%s diff %s --name-only", repoPath, repoPath, *revision)
	}
	out, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	split := func(c rune) bool {
		return c == '\n'
	}
	return strings.FieldsFunc(string(out), split), err
}
