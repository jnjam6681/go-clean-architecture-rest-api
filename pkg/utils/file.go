package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteToJsonFile(pathFile string, content interface{}) error {
	data, err := json.MarshalIndent(content, "", "\t")
	if err != nil {
		return fmt.Errorf("error marshaling JSON file: %s", err)
	}

	err = os.WriteFile(pathFile, data, 0644)
	if err != nil {
		return fmt.Errorf("error write file: %s", err)
	}
	return nil
}

func ReadJsonFile(path string, content interface{}) error {
	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading JSON file: %s", err)
	}

	if err := json.Unmarshal(jsonFile, &content); err != nil {
		return fmt.Errorf("error parsing JSON file: %s", err)
	}
	return nil
}
