package cl

import (
	"bump/pairs"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func ParseVerify(args []string) bool {
	for _, s := range args {
		if isNoVerify(s) {
			return true
		}
	}
	return false
}

func ParseImages(args []string) (pairs.PairCollection, error) {
	var pairs = pairs.PairCollection{}
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

func ParseAutoTags(args []string) (*string, *string, error) {
	err := errors.New("please set console and/or kyma tags in format '-c <console tag> -k <kyma tag>'")
	if len(args) != 2 && len(args) != 4 {
		return nil, nil, err
	}
	var kymaTag, consoleTag string
	if args[0] == "-c" {
		consoleTag = args[1]
	} else if args[0] == "-k" {
		kymaTag = args[1]
	} else {
		return nil, nil, err
	}

	if len(args) == 4 {
		if args[2] == "-c" {
			consoleTag = args[3]
		} else if args[2] == "-k" {
			kymaTag = args[3]
		} else {
			return nil, nil, err
		}
	}
	return &consoleTag, &kymaTag, nil
}

func ParseVerifyRevision(args []string) *string {
	if len(args) == 0 {
		return nil
	}

	return &args[0]
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
