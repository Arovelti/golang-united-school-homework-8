package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	Id        string = "id"
	Item             = "item"
	FileName         = "fileName"
	Operation        = "operation"
	ErrorMsg         = "-%s flag has to be specified"

	FindById = "findById"
	List     = "list"
	Add      = "add"
	Remove   = "remove"
)

type FileRecord struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Arguments map[string]string

func Perform(args Arguments, writer io.Writer) error {
	var result []byte
	var err error
	var fileRecords []FileRecord
	var operationValue = args[Operation]

	// check common arguments
	if args[FileName] == "" {
		return errors.New(fmt.Sprintf(ErrorMsg, FileName))
	}
	if operationValue == "" {
		return errors.New(fmt.Sprintf(ErrorMsg, Operation))
	}

	// read data from file
	fileRecords, err = getJsonDataFromFile(args[FileName])

	switch operationValue {
	case List:
		var data, _ = json.Marshal(fileRecords)
		result = data
		break
	case Add:
		result, err = addElementToFile(args[Item], args[FileName], fileRecords)
		break
	case FindById:
		result, err = findElementById(args[Id], fileRecords)
		break
	case Remove:
		result, err = removeElementById(args[Id], args[FileName], fileRecords)
		break
	default:
		return errors.New(fmt.Sprintf("Operation %s not allowed!", operationValue))
	}

	writer.Write(result)
	return err
}

func main() {
	err := Perform(parseArgs(), os.Stdout)

	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {
	var idF = flag.String(Id, "", "user id value")
	var operationF = flag.String(Operation, "", "add, remove, find by ID or list")
	var itemF = flag.String(Item, "", "item in json format with, id, email and age")
	var filenameF = flag.String(FileName, "", "Filepath to update")
	flag.Parse()
	return Arguments{Operation: *operationF, Id: *idF, Item: *itemF, FileName: *filenameF}
}
