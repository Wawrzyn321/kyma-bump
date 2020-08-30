package requirements

import (
	"bump/fileIO"
	"errors"
	"fmt"
	"os"
)

func getKymaPathVariableName() string {
	return "BUMP_KYMA_HOME"
}

func GetKymaPath() string {
	return os.Getenv(getKymaPathVariableName())
}

func CheckKymaPathRequirement() error {
	kymaPath := GetKymaPath()
	if kymaPath == "" {
		return errors.New(fmt.Sprintf("'%s' variable is not set. Please point it to your Kyma repo directory", getKymaPathVariableName()))
	} else if !fileIO.DirExists(kymaPath) {
		return errors.New(fmt.Sprintf("'%s' refers to non-existing directory. Please point it to your Kyma repo directory", getKymaPathVariableName()))
	}
	return nil
}

func PrintKymaPathRequirement() {
	fmt.Println("bump utility needs to know your Kyma path.")
	fmt.Printf("Please set %s variable to your Kyma repository path.\n", getKymaPathVariableName())
}
