package go_samples

import (
	"encoding/json"
	"fmt"
)

type Configuration struct {
	LogLevel int
}

func getConfiguration(filename string) (Configuration, error) {
	config := Configuration{}
	data, err := readFile(filename)
	if err != nil {
		return config, err
	} else {
		jsonErr := json.Unmarshal(data, &config)
		if jsonErr != nil {
			return config, jsonErr
		} else {
			return config, nil
		}
	}
}

func printConfiguration(config Configuration) {
	configJson, _ := json.MarshalIndent(config, "", "  ")
	fmt.Printf("Configuration:\n %s\n", string(configJson))
}
