package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadDirectory(directory string, format string) ([]string, error) {
	info, err := os.Stat(directory)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist")
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("path is not a directory")
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var specificFormatFiles []string

	fileFormat := "." + format

	for _, file := range files {
		name := file.Name()
		extend := filepath.Ext(name)
		if !file.IsDir() && extend == fileFormat {
			fileNameWithoutExt := strings.TrimSuffix(name, extend)
			specificFormatFiles = append(specificFormatFiles, fileNameWithoutExt)
		}
	}
	return specificFormatFiles, nil
}

func ReadCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var records [][]string

	records, err = reader.ReadAll()
	if err != nil {
		return records, err
	}
	return records, nil
}
