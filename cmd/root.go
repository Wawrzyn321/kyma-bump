package cmd

import (
	"bump/requirements"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		printBranding()
		fmt.Println("")

		requirements.PrintPathsRequirement()
		requirements.PrintDockerRequirement()

		fmt.Println("")
		fmt.Println("bump -h for more help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func printBranding() {
	fmt.Println(strings.Repeat("*", 50))
	fmt.Println("* *                  Mr. BUMP                  * *")
	fmt.Println(strings.Repeat("*", 50))
}