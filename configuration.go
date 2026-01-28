package go_samples

import (
	"encoding/json"
)

type Configuration struct {
	LogLevel int
}

func (c *Configuration) String() string {
	config, _ := json.MarshalIndent(c, "", "  ")
	return string(config)
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
