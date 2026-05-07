package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fileManager FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fileManager.InputFilePath)
	if err != nil {
		return nil, errors.New("Could not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, errors.New("Reading file content failed")
	}

	return lines, nil
}

func (fileManager FileManager) WriteResult(data any) error {
	file, err := os.Create(fileManager.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create file.")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Failed to encode data to JSON.")
	}

	return nil
}
