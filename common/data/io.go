package data

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(homeDir, ".mp", "config.json")
}

func GetConfig() (Config, error) {
	configFilePath := getConfigFilePath()

	file, err := os.Open(configFilePath)
	if err != nil {
		return getDefaultConfig(), err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return getDefaultConfig(), err
	}

	return config, nil
}

func SetConfig(config Config) error {
	configFilePath := getConfigFilePath()

	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		return err
	}

	return nil
}
