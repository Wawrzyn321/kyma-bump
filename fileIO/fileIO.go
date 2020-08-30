package fileIO

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func ReadFile(path string) ([]*string, error) {
	readFile, err := os.Open(path)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open file: %s", err))
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []*string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fileTextLines = append(fileTextLines, &line)
	}

	readFile.Close()

	return fileTextLines, nil
}

func WriteFile(path string, lines []*string) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New(fmt.Sprintf("failed to open file: %s", err))
	}
	defer file.Close()

	for _, line := range lines {
		file.WriteString(*line)
		file.WriteString("\n")
	}
	return nil
}
