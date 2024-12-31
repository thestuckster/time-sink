package internal

import (
	"encoding/json"
	"errors"
	"os"
)

func LoadConfiguration() TimeSinkConfig {
	path := getConfigFilePath()
	if !configFileExists(path) {
		file, err := os.Create(path)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		config := defaultConfig()
		writeConfigFile(config, path)

		return config
	}

	return loadConfigFile(path)
}

func SaveConfiguration(config TimeSinkConfig) {
	path := getConfigFilePath()
	writeConfigFile(config, path)
}

func getConfigFilePath() string {
	path, err := os.UserHomeDir()
	path += "/time-sink-config.json"
	if err != nil {
		panic(err)
	}

	return path
}

func configFileExists(path string) bool {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func defaultConfig() TimeSinkConfig {
	return TimeSinkConfig{
		Applications:  make([]string, 0),
		CheckInterval: "1 m",
	}
}

func writeConfigFile(config TimeSinkConfig, path string) {
	jsonData, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func loadConfigFile(path string) TimeSinkConfig {
	contents, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var config TimeSinkConfig
	err = json.Unmarshal(contents, &config)
	if err != nil {
		panic(err)
	}

	return config
}
