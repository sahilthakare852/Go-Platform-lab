package health

import (
	"encoding/json"
	"os"
)

func LoadConfig(filename string) (ServiceConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return ServiceConfig{}, err
	}
	var config ServiceConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return ServiceConfig{}, err
	}
	return config, nil
}
