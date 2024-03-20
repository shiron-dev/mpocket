package data

import (
	"encoding/json"
	"io/ioutil"
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

	jsonData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return getDefaultConfig(), err
	}

	var config Config
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return getDefaultConfig(), err
	}

	return config, nil
}

func SetConfig(config Config) error {
	configFilePath := getConfigFilePath()

	jsonData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFilePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
