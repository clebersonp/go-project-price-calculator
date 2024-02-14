package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
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

func WriteJSON(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file")
	}

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
