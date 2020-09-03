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

func VerifyImages(revision *string, m mappings.Mappings) error {
	filePaths, err := sysCommands.GetChangedFiles(requirements.GetKymaPath(), revision)
	if err != nil {
		return errors.New(fmt.Sprintf("Error during git diff: %s.", err))
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
	return nil
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
				fmt.Printf("WARN: %s: image %s not found in image registry. File left untouched.\n", mapping.Name, tag)
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

func Auto(m mappings.Mappings, consoleTag, kymaTag *string, noVerify bool) error {
	// it'd be better if we could return an array of string from getChangesIn... functions
	// BUT THERE ARE NO HASHSETS IN GO
	pairs := pairs.PairCollection{}
	err := getChangesInConsole(m, consoleTag, kymaTag, "master", pairs)
	if err != nil {
		return err
	}
	err = getChangesInKyma(consoleTag, "master", pairs)
	if err != nil {
		return err
	}

	if len(pairs) != 0 {
		fmt.Printf("Detected changes in ")
		// god forgive me for I have sinned
		// where is LINQ?
		index := 0
		var last string
		for name, _ := range pairs {
			index++
			if index != len(pairs) {
				fmt.Print(name)
				fmt.Print(", ")
			} else {
				last = name
			}
		}
		fmt.Print(last) // we don't want a ','
		fmt.Println(".")
		BumpImages(m, pairs, noVerify)
	} else {
		fmt.Println("No changes detected,.")
	}
	return nil
}

func getChangesInConsole(m mappings.Mappings, consoleTag, kymaTag *string, revision string, pairs pairs.PairCollection) error {
	path := requirements.GetConsolePath()
	lines, err := sysCommands.GetChangedFiles(path, &revision)
	if err != nil {
		return err
	}
	for _, diffLine := range lines {
		for _, mapping := range m {
			if strings.HasPrefix(diffLine, "tests") && kymaTag != nil {
				pairs[mapping.Name] = *kymaTag
			} else if strings.HasPrefix(diffLine, mapping.Name + "/") && consoleTag != nil {
				pairs[mapping.Name] = *consoleTag
			}
		}
	}
	return nil
}

func getChangesInKyma(kymaTag *string, revision string, pairs pairs.PairCollection) error {
	if kymaTag == nil {
		return nil
	}
	path := requirements.GetKymaPath()
	lines, err := sysCommands.GetChangedFiles(path, &revision)
	if err != nil {
		return err
	}
	for _, diffLine := range lines {
		if strings.HasPrefix(diffLine, "components/console-backend-service") {
			pairs["console-backend-service"] = *kymaTag
		} else if strings.HasPrefix(diffLine, "tests/console-backend-service") {
			pairs["console-backend-service-test"] = *kymaTag
		}
	}
	return nil
}

func extractTagFromLine(line *string) string {
	return (*line)[strings.Index(*line, ": ") + 2:]
}

func findLineNo(lines []*string, yamlPath string) (*int, error) {
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
