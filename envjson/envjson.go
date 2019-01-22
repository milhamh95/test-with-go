package envjson

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var envkey map[string]string

func Setup() error {

	cwd, err := os.Getwd()
	if err != nil {
		return errors.New("Failed to get current working directory")
	}

	rawJson, err := ioutil.ReadFile(filepath.Join(cwd, "env.json"))
	if err != nil {
		return errors.New("Failed to read env.json in the directory")
	}

	err = json.Unmarshal(rawJson, &envkey)
	if err != nil {
		return errors.New("Failed to unmarshall env.json")
	}

	for key, value := range envkey {
		os.Setenv(key, value)
	}

	return nil
}
