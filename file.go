package main

import (
	"encoding/json"
	"os"
)

func writeDataToFile(fileName string, data []byte) ([]byte, error) {
	os.Remove(fileName)
	var err = os.WriteFile(fileName, data, 0644)
	return data, err
}

func getJsonDataFromFile(fileName string) ([]FileRecord, error) {
	var fileRecords []FileRecord
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &fileRecords)
	if err == nil {
		return fileRecords, err
	}
	return nil, err
}
