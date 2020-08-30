package cl

import (
	"bump/mappings"
	"bump/requirements"
	"fmt"
	"strings"
)

func PrintHelp() {
	printCommandsHelp()
	requirements.PrintKymaPathRequirement()
	requirements.PrintDockerRequirement()
}

func printCommandsHelp() {
	fmt.Println("  Commands:")
	fmt.Println("    check-files")
	fmt.Println("      Checks if files exist and their paths match.")
	fmt.Println("    img")
	fmt.Println("      Updates tags of images. Usage:")
	fmt.Println("        bump img <tag1> <...images> <tag2> <...images>")
	fmt.Println("      You can use either commit hash or PR tag. In former case, at least 8 characters of tag is required.")
	fmt.Println("      Add --no-verify or -f flag to disable image check.")
	fmt.Println("    help, -h")
	fmt.Println("      Prints help.")
	fmt.Println("    list, -l")
	fmt.Println("      Lists all supported images, along with their Aliases.")
	fmt.Println("    verify")
	fmt.Println("      Verifies changed images in repo, comparing current HEAD. Branch defaults to master TODO")
	fmt.Println("")
}

func List(m mappings.Mappings) {
	fmt.Println("Supported images:")
	for _, mapping := range m {
		fmt.Printf("%30s     %s\n", mapping.Name, strings.Join(mapping.Aliases, ", "))
	}
	fmt.Println("Yup, I dunno how to format in Go.")
}
