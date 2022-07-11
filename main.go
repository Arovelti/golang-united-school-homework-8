package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	Id        = "id"
	Item      = "item"
	FileName  = "fileName"
	Operation = "operation"
	ErrorMsg  = "-%s flag has to be specified"

	FindById = "findById"
	List     = "list"
	Add      = "add"
	Remove   = "remove"
)

type Arguments map[string]string

type FileRecord struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func Perform(args Arguments, writer io.Writer) error {
	var result []byte
	var err error
	var FileRecords []FileRecord
	var operationValue = args[Operation]

	if args[FileName] == "" {
		return fmt.Errorf(ErrorMsg, FileName)
	}
	if operationValue == "" {
		return fmt.Errorf(ErrorMsg, Operation)
	}

	fileRecords, err := getJsonDataFromFile(args[FileName])

	switch operationValue {
	case List:
		var data, err = json.Marshal(FileRecords)
		if err != nil {
			return err
		}
		result = data
	case Add:
		result, err = addElementToFile(args[Item], args[FileName], fileRecords)
	case FindById:
		result, err = findElementById(args[Id], FileRecords)
	case Remove:
		result, err = removeElementById(args[Id], args[FileName], fileRecords)
	default:
		return fmt.Errorf("operation %s not allowed", operationValue)
	}

	writer.Write(result)
	return err
}

func parseArgs() Arguments {
	var idF = flag.String(Id, "", "user id value")
	var operationF = flag.String(Operation, "", "add, remove, find by ID or list")
	var itemF = flag.String(Item, "", "item in json format")
	var filenameF = flag.String(FileName, "", "Filepath to update")
	flag.Parse()
	return Arguments{Operation: *operationF, Id: *idF, Item: *itemF, FileName: *filenameF}
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
