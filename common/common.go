package common

import (
	"bump/fileIO"
	"bump/model"
	"bump/requirements"
	"bump/sysCommands"
	"errors"
	"fmt"
	"path"
	"strings"
)

func BumpImages(m model.Mappings, pairs model.PairCollection, noVerify bool) {
	pairs, errs := pairs.Dealiasize(m)
	for _, err := range errs {
		fmt.Printf("WARN %s\n", err)
	}
	for image, tag := range pairs {
		mapping := m.Find(image)
		if mapping == nil {
			fmt.Printf("Cannot find mapping for %s.\n", image)
			continue
		}

		fmt.Printf("Updating %s tag to %s... ", mapping.Name, tag)

		if !noVerify {
			ok := sysCommands.CheckImageExists(mapping.RegistryUrl, tag)
			if !ok {
				fmt.Printf("WARN: image %s not found in image registry. File left untouched.\n", tag)
				break
			}
		}

		basedir := requirements.GetKymaPath()
		var filePath = path.Join(basedir, mapping.FilePath)
		lines, err := fileIO.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error during opening %s: %s.\n", image , err)
			break
		}
		i, err := FindLineNo(lines, mapping.YamlPath)
		if err != nil {
			fmt.Println(err)
			break
		}

		indent := strings.Repeat(" ", 2 * strings.Count(mapping.YamlPath, "."))
		property := mapping.YamlPath[strings.LastIndex(mapping.YamlPath, ".") + 1:]
		lookup := fmt.Sprintf("%s%s:", indent, property)
		var newLine = fmt.Sprintf("%s %s", lookup, tag)
		lines[*i] = &newLine
		err = fileIO.WriteFile(filePath, lines)
		if err != nil {
			fmt.Printf("Error during writing %s: %s.\n", image , err)
			break
		}

		fmt.Println("ok!")
	}
}

func FindLineNo(lines []*string, yamlPath string) (*int, error) {
	return findLineNoSub(lines, strings.Split(yamlPath, "."), 0, 0)
}

func findLineNoSub(lines []*string, subPaths []string, subPathIndex int, start int) (*int, error) {
	if subPathIndex > len(subPaths) {
		return nil, errors.New("path out of range")
	}
	lookup := fmt.Sprintf("%s%s:", strings.Repeat(" ", 2 * subPathIndex), subPaths[subPathIndex])

	var newStart *int
	for i := start; i < len(lines); i++ {
		var line = lines[i]
		if strings.HasPrefix(*line, lookup) {
			var ii = i
			newStart = &ii
			break
		}
	}
	if newStart == nil {
		return nil, errors.New("Could not find TODO better msg.")
	}

	if subPathIndex == len(subPaths) - 1 {
		return newStart, nil
	}
	return findLineNoSub(lines, subPaths, subPathIndex + 1, *newStart)
}
