package config

import (
	"os"
	"encoding/json"
)

func LoadFromFile(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
