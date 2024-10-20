package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadDirectory(directory string) ([]string, error) {
	var specificFormatFiles []string

	info, err := os.Stat(directory)
	if os.IsNotExist(err) {
		return specificFormatFiles, fmt.Errorf("directory does not exist")
	}

	if !info.IsDir() {
		return specificFormatFiles, fmt.Errorf("path is not a directory")
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return specificFormatFiles, err
	}

	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && filepath.Ext(name) == ".csv" {
			fileNameWithoutExt := strings.TrimSuffix(name, filepath.Ext(name))
			specificFormatFiles = append(specificFormatFiles, fileNameWithoutExt)
		}
	}
	return specificFormatFiles, nil
}

func ReadCSV(path string) ([][]string, error) {
	var records [][]string

	file, err := os.Open(path)
	if err != nil {
		return records, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err = reader.ReadAll()
	if err != nil {
		return records, err
	}
	return records, nil
}
