package main

import (
	"encoding/json"
	"fmt"
)

func findElementById(id string, fileRecords []FileRecord) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf(ErrorMsg, Id)
	}

	for i := 0; i < len(fileRecords); i++ {
		if fileRecords[i].Id == id {
			return json.Marshal(fileRecords[i])
		}
	}
	return []byte(""), nil
}

func removeElementById(id string, fileName string, fileRecords []FileRecord) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf(ErrorMsg, Id)
	}

	modifiedRecords := make([]FileRecord, 0)
	for i := 0; i < len(fileRecords); i++ {
		if fileRecords[i].Id != id {
			modifiedRecords = append(modifiedRecords, fileRecords[i])
		}
	}

	if len(modifiedRecords) != len(fileRecords) {
		data, _ := json.Marshal(modifiedRecords)
		return writeDataToFile(fileName, data)
	}
	return []byte(fmt.Sprintf("Item with id %s not found", id)), nil
}

func addElementToFile(items string, fileName string, fileRecords []FileRecord) ([]byte, error) {
	if items == "" {
		return nil, fmt.Errorf(ErrorMsg, Item)
	}

	var itemsToAdd FileRecord
	var err = json.Unmarshal([]byte(items), &itemsToAdd)
	if err != nil {
		return nil, err
	}
	modifiedRecords := make([]FileRecord, 0)
	var id = itemsToAdd.Id
	for i := 0; i < len(fileRecords); i++ {
		if fileRecords[i].Id == id {
			return []byte(fmt.Sprintf("Item with id %s already exists", id)), nil
		}
	}
	modifiedRecords = append(modifiedRecords, fileRecords...)
	modifiedRecords = append(modifiedRecords, itemsToAdd)
	data, _ := json.Marshal(modifiedRecords)
	return writeDataToFile(fileName, data)
}
