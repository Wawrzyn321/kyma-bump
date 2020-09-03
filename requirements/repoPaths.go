package requirements

import (
	"bump/fileIO"
	"errors"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"os"
)

func getKymaPathVariableName() string {
	return "BUMP_KYMA_HOME"
}

func getConsolePathVariableName() string {
	return "BUMP_CONSOLE_HOME"
}

func GetKymaPath() string {
	return os.Getenv(getKymaPathVariableName())
}

func GetConsolePath() string {
	return os.Getenv(getConsolePathVariableName())
}

func CheckKymaPathRequirement() error {
	return checkPathRequirement(getKymaPathVariableName(), "Kyma")
}
func CheckConsolePathRequirement() error {
	return checkPathRequirement(getConsolePathVariableName(), "Console")
}

func checkPathRequirement(variable string, repoName string) error {
	path := os.Getenv(variable)
	if path == "" {
		return errors.New(fmt.Sprintf("'%s' variable is not set. Please point it to your %s repo directory", Bold(variable), repoName))
	} else if !fileIO.DirExists(path) {
		return errors.New(fmt.Sprintf("'%s' refers to non-existing directory. Please point it to your %s repo directory", Bold(variable), repoName))
	}
	return nil
}

func PrintPathsRequirement() {
	fmt.Println("bump utility needs to know your Kyma and Console paths.")
	fmt.Printf("Please set %s and %s variables to your Kyma and Console repository paths.\n", Bold(getKymaPathVariableName()), Bold(getConsolePathVariableName()))
}
