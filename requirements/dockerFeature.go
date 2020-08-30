package requirements

import (
	"errors"
	"fmt"
	"os"
)

func getDockerFeatureVariableName() string {
	return "DOCKER_CLI_EXPERIMENTAL"
}

func GetDockerFeatureStatus() string {
	return os.Getenv(getDockerFeatureVariableName())
}

func CheckDockerFeatureRequirement() error {
	dockerExperimental := GetDockerFeatureStatus()
	if dockerExperimental != "enabled" {
		return errors.New(fmt.Sprintf("to enable 'verify' command, set %s=enabled", getDockerFeatureVariableName()))
	}
	return nil
}


func PrintDockerRequirement() {
	fmt.Printf("To verify your images, you have to set your %s variable to 'enabled'\n", getDockerFeatureVariableName())
}
