package auction

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func getArrayFromFile(filePath string) ([]map[string]interface{}, error) {
	// Opening and reading the file into a byte array
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("could not open file")
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("could not read file")
		return nil, err
	}

	// Unmarshalling into a an array of maps/objects which will have interface (any) value types
	var itemArray []map[string]interface{}
	err = json.Unmarshal(byteValue, &itemArray)

	return itemArray, nil
}