package requirements

import (
	"errors"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"os"
)

func getDockerFeatureVariableName() string {
	return "DOCKER_CLI_EXPERIMENTAL"
}

func GetDockerFeatureStatus() string {
	return os.Getenv(getDockerFeatureVariableName())
}

func CheckDockerFeature() error {
	dockerExperimental := GetDockerFeatureStatus()
	if dockerExperimental != "enabled" {
		return errors.New(fmt.Sprintf("to enable 'verify' command, set %s=enabled", Bold(getDockerFeatureVariableName())))
	}
	return nil
}


func PrintDockerRequirement() {
	fmt.Printf("To verify your images, you have to set your %s variable to 'enabled'.\n", Bold(getDockerFeatureVariableName()))
}
