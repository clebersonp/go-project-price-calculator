package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string `json:"input_file_path"`
	OutputFilePath string `json:"output_file_path"`
}

// New - in this constructor we don't need to return a pointer because it's so smaller and doesn't manner
func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

// ReadLines - it does not need to receive a file manager pointer because we don't need to change its values, just read
func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	// get lines values
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		_ = file.Close() // ignore error
		return nil, errors.New("failed to load data from file")
	}

	_ = file.Close() // ignore error
	return lines, nil
}

// WriteResult - it does not need to receive a file manager pointer because we don't need to change its values, just read
func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}

	// simulate delay to work with goroutines
	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		_ = file.Close()
		return errors.New("failed to convert data to JSON format")
	}

	// another approach to write JSON to a file
	//marshal, err := json.MarshalIndent(data, "", "  ")
	//
	//if err != nil {
	//	_ = file.Close()
	//	return errors.New("failed to convert data to JSON format")
	//}
	//
	//_, err = file.Write(marshal)
	//
	//if err != nil {
	//	_ = file.Close()
	//	return errors.New("failed to write data into a file")
	//}

	_ = file.Close()
	return nil
}
