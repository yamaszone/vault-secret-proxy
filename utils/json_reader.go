package utils

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"os"
)

type KeyValueConfig map[string]interface{}

// ReadJsonFile parses JSON file containing kv-pairs and builds a map
func ReadJsonFile(filename string) (KeyValueConfig, error) {

	jsonFile, err := os.Open("/tmp/kv-data.json")
	if err != nil {
		log.Printf("Error opening kv-data.json: %s", err)
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Printf("Error reading kv-data.json: %s", err)
	}

	var response KeyValueConfig
	err = json.Unmarshal(jsonBytes, &response)
	if err != nil {
		log.Printf("Error parsing JSON: %s", err)
	}
	//fmt.Println(response)

	return response, nil
}
