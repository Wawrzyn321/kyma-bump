package main

import (
	"bump/bump"
	"bump/cl"
	"bump/mappings"
	"bump/requirements"
	"fmt"
	"os"
)

//todo change direction of <commit> <...images> to <...images <commit> ?
//todo requirements checks in switch are effing ugly
//todo pushing around mappingPresets as m is boring

func main() {
	const Success = 0
	const RequirementsNotMet = 1
	const ParseError = 2
	const VerificationFailed = 3
	const UnrecognizedCommand = 4
	const GenericCommandFailure = 5

	mappingPresets := mappings.GetMappingPresets()

	if len(os.Args) < 2 {
		cl.PrintHelp()
		os.Exit(Success)
	}

	switch os.Args[1] {
	case "auto":
		err := requirements.CheckKymaPathRequirement()
		if err != nil {
			fmt.Println(err)
			os.Exit(RequirementsNotMet)
		}
		err = requirements.CheckConsolePathRequirement()
		if err != nil {
			fmt.Println(err)
			os.Exit(RequirementsNotMet)
		}
		args := os.Args[2:]
		consoleTag, kymaTag, err := cl.ParseAutoTags(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(ParseError)
		}
		noVerify := cl.ParseVerify(args)
		err = bump.Auto(mappingPresets, consoleTag, kymaTag, noVerify)
		if err != nil {
			fmt.Println(err)
		}
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
		noVerify := cl.ParseVerify(args)
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
		revision := cl.ParseVerifyRevision(os.Args[2:])
		err = bump.VerifyImages(revision, mappingPresets)
		if err != nil {
			fmt.Println(err)
			os.Exit(GenericCommandFailure)
		}
		break
	default:
		fmt.Printf("Unrecognized command %s.\n\n", os.Args[1])
		cl.PrintHelp()
		os.Exit(UnrecognizedCommand)
	}
}
