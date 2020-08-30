package sysCommands

import (
	"fmt"
	"os/exec"
)

func CheckImageExists(registryUrl, tag string) bool {
	cmdStr := fmt.Sprintf("docker manifest inspect %s:%s", registryUrl, tag)
	_, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	return err == nil
}

