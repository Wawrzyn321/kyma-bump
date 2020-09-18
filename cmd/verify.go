package cmd

import (
	"bump/common"
	"bump/fileIO"
	"bump/model"
	"bump/requirements"
	"bump/sysCommands"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"path"
	"strings"
)

type verifyOptions struct {
	branch string
}

func init() {
	rootCmd.AddCommand(verifyCmd())
}

func addVerifyCmdFlags(cmd *cobra.Command, options *verifyOptions) {
	cmd.Flags().StringVarP(&options.branch, "branch", "b", "master", "branch to diff images by. Defaults to master.")
}

func verifyCmd() *cobra.Command {
	options := verifyOptions{}
	var cmd = &cobra.Command{
		Use:   "verify",
		Short: "Verifies changed images in repo",
		Long:  `Looks for changes in Kyma configuration (against chosen branch) and checks if changed images actually
exist in image registry.
Requirements: Kyma path, Docker experimental features.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := requirements.Check(requirements.CheckKymaPath, requirements.CheckConsolePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = verifyImages(options.branch, model.GetMappingPresets())
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	addVerifyCmdFlags(cmd, &options)
	return cmd
}

func verifyImages(branch string, m model.Mappings) error {
	filePaths, err := sysCommands.GetChangedFiles(requirements.GetKymaPath(), branch)
	if err != nil {
		return errors.New(fmt.Sprintf("Error during git diff: %s.", err))
	}
	fmt.Printf("Found %d changed files.\n", len(filePaths))

	kymaPath := requirements.GetKymaPath()
	for _, mapping := range m {
		found := false
		for _, path := range filePaths {
			if strings.HasPrefix(path, mapping.FilePath) {
				found = true
			}
		}
		if !found {
			continue
		}

		fmt.Printf("Checking %s... ", mapping.Name)
		lines, err := fileIO.ReadFile(path.Join(kymaPath, mapping.FilePath))
		if err != nil {
			fmt.Println("file not found!")
			continue
		}
		i, err := common.FindLineNo(lines, mapping.YamlPath)
		if err != nil {
			fmt.Println("property not found!")
			continue
		}
		tag := extractTagFromLine(lines[*i])
		exists := sysCommands.CheckImageExists(mapping.RegistryUrl, tag)
		if exists {
			fmt.Println("OK!")
		} else {
			fmt.Println("not found in image registry.")
		}
	}
	return nil
}

func extractTagFromLine(line *string) string {
	return (*line)[strings.Index(*line, ": ") + 2:]
}