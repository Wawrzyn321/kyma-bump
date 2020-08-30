package main

import (
	"bump/bump"
	"bump/cl"
	"bump/mappings"
	"bump/requirements"
	"fmt"
	"os"
)

//todo branch diffing in verify
//todo change direction of <commit> <...images> to <...images <commit> ?

func main() {
	const Success = 0
	const RequirementsNotMet = 1
	const ParseError = 2
	const VerificationFailed = 3
	const UnrecognizedCommand = 4

	mappingPresets := mappings.GetMappingPresets()

	if len(os.Args) < 2 {
		cl.PrintHelp()
		os.Exit(Success)
	}

	switch os.Args[1] {
		case "img":
			err := requirements.CheckKymaPathRequirement()
			if err != nil {
				fmt.Println(err)
				os.Exit(RequirementsNotMet)
			}
			args := os.Args[2:]
			pairs, err := cl.ParseImages(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(ParseError)
			}
			noVerify := cl.ShouldNoVerify(args)
			bump.BumpImages(mappingPresets, pairs, noVerify)
			break
		case "help":
			cl.PrintHelp()
			break
		case "-h":
			cl.PrintHelp()
			break
		case "list":
			cl.List(mappingPresets)
			break
		case "verify":
			err := requirements.CheckKymaPathRequirement()
			if err != nil {
				fmt.Println(err)
				os.Exit(RequirementsNotMet)
			}
			err = requirements.CheckDockerFeatureRequirement()
			if err != nil {
				fmt.Println(err)
				os.Exit(RequirementsNotMet)
			}

			bump.VerifyImages("HEAD" , mappingPresets)
			break
		case "check-files":
			err := requirements.CheckKymaPathRequirement()
			if err != nil {
				fmt.Println(err)
				os.Exit(RequirementsNotMet)
			}
			ok := bump.VerifyFiles(mappingPresets)
			if ok {
				fmt.Println("No problems found.")
			} else {
				os.Exit(VerificationFailed)
			}
			break
		default:
			fmt.Printf("Unrecognized command %s.\n\n", os.Args[1])
			cl.PrintHelp()
			os.Exit(UnrecognizedCommand)
	}
}
