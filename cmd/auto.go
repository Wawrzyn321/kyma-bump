package cmd

import (
	"bump/common"
	"bump/model"
	"bump/requirements"
	"bump/sysCommands"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

type ChangesList = map[string]bool // hAshSEt iN gOLanG

type autoOptions struct {
	consoleTag string
	kymaTag string
	branch string
	noVerify bool
}

func init() {
	rootCmd.AddCommand(autoCmd())
}

func addAutoCmdFlags(cmd *cobra.Command, options *autoOptions) {
	cmd.Flags().StringVarP(&options.consoleTag, "console-tag", "c", "", "tag for Console repo images")
	cmd.Flags().StringVarP(&options.kymaTag, "kyma-tag", "k", "", "tag for Kyma repo images")
	cmd.Flags().StringVarP(&options.branch, "branch", "b", "", "branch to diff images by. Defaults to master.")
	cmd.Flags().BoolVarP(&options.noVerify, "no-verify", "f", false, "disable image check")
}

func autoCmd() *cobra.Command {
	options := autoOptions{}
	var cmd = &cobra.Command{
		Use:   "auto",
		Short: "Set images based on git status",
		Long:  `Diffs current state of Kyma and Console and bumps images
At least one (-c, -k) is required.
You can use either commit hash or PR tag. In former case, at least 8 characters of tag is required
Requirements: Kyma and Console repo paths.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := requirements.Check(requirements.CheckKymaPath, requirements.CheckConsolePath)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			if options.kymaTag == "" && options.consoleTag == "" {
				return errors.New("neither Kyma nor Console tag specified")
			}
			err = auto(model.GetMappingPresets(), options)
			if err != nil {
				fmt.Println(err)
			}
			return nil
		},
	}
	addAutoCmdFlags(cmd, &options)
	return cmd
}

func auto(presets model.Mappings, options autoOptions) error {
	if options.consoleTag != "" {
		err := autoConsole(presets, options)
		if err != nil {
			return err
		}
	}
	if options.kymaTag != "" {
		err := autoKyma(presets, options)
		if err != nil {
			return err
		}
	}

	return nil
}

func autoConsole(presets model.Mappings, options autoOptions) error {
	fmt.Print("Checking changes in Console... ")
	list, err := getChangesInConsole(presets, options.branch)
	if err != nil {
		return err
	}
	fmt.Printf("detected changes in ")
	printDetectedChanges(list)
	pairs := makePairs(presets, list, options.consoleTag)
	common.BumpImages(presets, pairs, options.noVerify)
	return nil
}

func autoKyma(presets model.Mappings, options autoOptions) error {
	fmt.Print("Checking changes in Kyma... ")
	list, err := getChangesInKyma(presets, options.branch)
	if err != nil {
		return err
	}
	fmt.Printf("detected changes in ")
	printDetectedChanges(list)
	pairs := makePairs(presets, list, options.kymaTag)
	common.BumpImages(presets, pairs, options.noVerify)
	return nil
}

func getChangesInConsole(m model.Mappings, branch string) (ChangesList, error) {
	changes := make(ChangesList)

	path := requirements.GetConsolePath()
	lines, err := sysCommands.GetChangedFiles(path, branch)
	if err != nil {
		return nil, err
	}
	for _, diffLine := range lines {
		for _, mapping := range m {
			if strings.HasPrefix(diffLine, "tests") {
				changes[mapping.Name] = true
			} else if strings.HasPrefix(diffLine, mapping.Dir) {
				changes[mapping.Name] = true
			}
		}
	}
	return changes, nil
}

func getChangesInKyma(m model.Mappings, branch string) (ChangesList, error) {
	changes := make(ChangesList)

	path := requirements.GetKymaPath()
	lines, err := sysCommands.GetChangedFiles(path, branch)
	if err != nil {
		return nil, err
	}
	for _, diffLine := range lines {
		for _, mapping := range m {
			if strings.HasPrefix(diffLine, mapping.Dir) {
				changes[mapping.Name] = true
			}
		}
	}
	return changes, nil
}

func printDetectedChanges(list ChangesList) {
	// god forgive me for I have sinned
	// where is LINQ?
	index := 0
	var last string
	for name, _ := range list {
		index++
		if index != len(list) {
			fmt.Print(name)
			fmt.Print(", ")
		} else {
			last = name
		}
	}
	fmt.Print(last) // we don't want a ',' at the end
	fmt.Println(".")
}

func makePairs(presets model.Mappings, list ChangesList, tag string) model.PairCollection {
	pairs := model.PairCollection{}
	for mappingName, _ := range list {
		for _, mapping := range presets {
			if mappingName == mapping.Name {
				pairs[mappingName] = tag
			}
		}
	}
	return pairs
}

