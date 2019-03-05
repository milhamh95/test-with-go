package envjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var envkey map[string]string

func Setup(filePath string) error {

	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close()
	if err != nil {
		return errors.New("failed to open json file: " + err.Error())
	}

	rawJson, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to read env.json in the directory: " + err.Error())
	}

	err = json.Unmarshal(rawJson, &envkey)
	if err != nil {
		return errors.New("Failed to unmarshall env.json: " + err.Error())
	}

	for key, value := range envkey {
		os.Setenv(key, value)
	}

	return nil
}
