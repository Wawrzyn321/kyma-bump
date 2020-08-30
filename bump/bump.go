package bump

import (
	"bump/fileIO"
	"bump/mappings"
	"bump/pairs"
	"bump/requirements"
	"bump/sysCommands"
	"errors"
	"fmt"
	"path"
	"strings"
)

func VerifyImages(revision string, m mappings.Mappings) {
	filePaths, err := sysCommands.GetChangedFiles(revision)
	if err != nil {
		fmt.Printf("Error during git diff: %s.", err)
		return
	}
	fmt.Printf("Found %d changed files.\n", len(filePaths))

	kymaPath := requirements.GetKymaPath()
	for _, mapping := range m {
		fmt.Println(mapping.Name)
		found := false
		for _, path := range filePaths {
			if path == mapping.FilePath {
				found = true
			}
		}
		fmt.Println(found)
		if !found {
			continue
		}

		fmt.Printf("Checking %s... ", mapping.Name)
		lines, err := fileIO.ReadFile(path.Join(kymaPath, mapping.FilePath))
		if err != nil {
			fmt.Println("file not found!")
			continue
		}
		i, err := findLineNo(lines, mapping.YamlPath)
		if err != nil {
			fmt.Println("property not found!")
			continue
		}
		tag := extractTagFromLine(lines[*i])
		exists := sysCommands.CheckImageExists(mapping.RegistryUrl, tag)
		if exists {
			fmt.Println("OK!")
		} else {
			fmt.Println("not found.")
		}
	}
}

func VerifyFiles(mappings mappings.Mappings) bool {
	ok := true
	kymaPath := requirements.GetKymaPath()
	for _, mapping := range mappings {
		filePath := path.Join(kymaPath, mapping.FilePath)
		if !fileIO.FileExists(filePath) {
			fmt.Printf("%s: not found file %s\n", mapping.Name, filePath)
			ok = false
			continue
		}

		lines, err := fileIO.ReadFile(filePath)
		if err != nil {
			fmt.Printf("%s: cannot open file %s: %s\n", mapping.Name, filePath, err)
			ok = false
			continue
		}
		_, err = findLineNo(lines, mapping.YamlPath)

		if err != nil {
			fmt.Printf("%s: Could not find tag by yaml path: %s (%s)\n", mapping.Name, mapping.YamlPath, mapping.FilePath)
			ok = false
		}
	}
	return ok
}

func BumpImages(m mappings.Mappings, pairs pairs.PairCollection, noVerify bool) {
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

		if !noVerify {
			ok := sysCommands.CheckImageExists(mapping.RegistryUrl, tag)
			if !ok {
				fmt.Printf("WARN: %s: image  %s not found in image registry. File left untouched.\n", mapping.Name, tag)
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
		i, err := findLineNo(lines, mapping.YamlPath)
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

		fmt.Printf("Updated %s tag to %s.\n", mapping.Name, tag)
	}
}

func extractTagFromLine(line *string) string {
	return (*line)[strings.Index(*line, ": ") + 2:]
}

func findLineNo(lines []*string, yamlPath string) (*int, error) {
	return _findLineNo(lines, strings.Split(yamlPath, "."), 0, 0)
}

func _findLineNo(lines []*string, subPaths []string, subPathIndex int, start int) (*int, error) {
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
	return _findLineNo(lines, subPaths, subPathIndex + 1, *newStart)
}
