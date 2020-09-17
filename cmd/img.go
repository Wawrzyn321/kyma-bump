package cmd

import (
	"bump/common"
	"bump/model"
	"bump/requirements"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

type imgOptions struct {
	noVerify bool
}

func init() {
	rootCmd.AddCommand(imgCmd())
}

func addImgCmdFlags(cmd *cobra.Command, options *imgOptions) {
	cmd.Flags().BoolVarP(&options.noVerify, "no-verify", "f", false, "")
}

func imgCmd() *cobra.Command {
	options := imgOptions{}
	var cmd = &cobra.Command{
		Use:   "img",
		Short: "Updates tags of images.",
		Long:  `You can use either commit hash or PR tag. In former case, at least 8 characters of tag is required.
The images will be verified against eu.gcr.io/kyma-project registry.
Requirements: Kyma repo path, Docker experimental features.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := requirements.CheckKymaPathRequirement()
			if err != nil {
				fmt.Println(err)
				return nil
			}
			err = requirements.CheckDockerFeatureRequirement()
			if err != nil {
				fmt.Println(err)
				return nil
			}

			pairs, err := parseImages(args)
			if err != nil {
				return err
			}

			common.BumpImages(model.GetMappingPresets(), pairs, options.noVerify)
			return nil
		},
	}
	cmd.SetHelpFunc(func(*cobra.Command, []string) {
		fmt.Println(`
Usage:
	common img <tag1> <...images> <tag2> <...images>
		`)
		fmt.Println(cmd.Long)
		fmt.Println(`Flags:
	  -h, --help        help for img
	  -f, --no-verify   disable image check - useful when your image is not yet built,
			    or you are using an image from custom registry.
	`)
	})
	addImgCmdFlags(cmd, &options)
	return cmd
}

func parseImages(args []string) (model.PairCollection, error) {
	var pairs = model.PairCollection{}
	var currentTag *string = nil
	for _, s := range args {
		if isNoVerify(s) {
			continue
		}
		if isTag(s) {
			tag := s
			if !isPRTag(s) {
				tag = s[:8] // only 8 characters required
			}
			currentTag = &tag
		} else {
			if currentTag == nil {
				return pairs, errors.New(fmt.Sprintf("Tag for image '%s' not set.\n", s))
			} else {
				pairs[s] = *currentTag
			}
		}
	}
	return pairs, nil
}

func isTag(s string) bool {
	match, _ := regexp.MatchString("^(([a-f]|[0-9]){8,})|(PR-\\d+)$", s)
	return match
}

func isPRTag(s string) bool {
	return strings.HasPrefix(s, "PR-")
}

func isNoVerify(s string) bool {
	return s == "--no-verify" || s == "-f"
}
