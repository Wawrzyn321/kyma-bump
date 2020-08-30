package cl

import (
	"bump/pairs"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func ShouldNoVerify(args []string) bool {
	for _, s := range args {
		if s == "--no-verify" || s== "-f" {
			return true
		}
	}
	return false
}

func ParseImages(args []string) (pairs.PairCollection, error) {
	var pairs = pairs.PairCollection{}
	var currentTag *string = nil
	for _, s := range args {
		if s == "--no-verify" {
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