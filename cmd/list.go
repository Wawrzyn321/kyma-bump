package cmd

import (
	"bump/model"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all supported images, along with their aliases",
	Long:  `You can add your own aliases in "mappingPresets.go"`,
	Run: func(cmd *cobra.Command, args []string) {
		list(model.GetMappingPresets())
	},
}

func list(m model.Mappings) {
	fmt.Println("Supported images:")
	for _, mapping := range m {
		fmt.Printf("%30s     %s\n", mapping.Name, strings.Join(mapping.Aliases, ", "))
	}
	fmt.Println("Yup, I dunno how to format in Go.")
}
