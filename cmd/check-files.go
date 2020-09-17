package cmd

import (
	"bump/common"
	"bump/fileIO"
	"bump/model"
	"bump/requirements"
	"fmt"
	"github.com/spf13/cobra"
	"path"
)

func init() {
	rootCmd.AddCommand(checkFilesCmd)
}

var checkFilesCmd = &cobra.Command{
	Use:    "check-files",
	Long:   `Checks if YAML configuration files exist and their tag variable paths match.
Requirements: Kyma repo path.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		err := requirements.CheckKymaPathRequirement()
		if err != nil {
			fmt.Println("Requirement not met")
			return
		}
		ok := verifyFiles(model.GetMappingPresets())
		if ok {
			fmt.Println("No problems found.")
		}
	},
}

func verifyFiles(mappings model.Mappings) bool {
	ok := true
	kymaPath := requirements.GetKymaPath()
	for _, mapping := range mappings {
		filePath := path.Join(kymaPath, mapping.FilePath)
		if !fileIO.FileExists(filePath) {
			fmt.Printf("%s: not found file %s\n", mapping.Name, filePath)
			ok = false
			continue
		}

		lines, err := fileIO.ReadFile(filePath)
		if err != nil {
			fmt.Printf("%s: cannot open file %s: %s\n", mapping.Name, filePath, err)
			ok = false
			continue
		}
		_, err = common.FindLineNo(lines, mapping.YamlPath)

		if err != nil {
			fmt.Printf("%s: Could not find tag by yaml path: %s (%s)\n", mapping.Name, mapping.YamlPath, mapping.FilePath)
			ok = false
		}
	}
	return ok
}

