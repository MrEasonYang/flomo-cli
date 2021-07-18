package application

import (
	"io/ioutil"
	"encoding/json"
	"os"
)

const CONFIG_FILE_NAME = ".flomo-cli.config"

// SaveConfig is the way to write flomo-cli config.
func SaveConfig(config FlomoConfig) {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(ParseFilePath(CONFIG_FILE_NAME), data, 0600)
	if err != nil {
		panic(err)
	}
}

// GetConfig is the way to read flomo-cli config.
func GetConfig() FlomoConfig {
	file, err := os.Open(ParseFilePath(CONFIG_FILE_NAME))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var config FlomoConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}

	return config
}

// FlomoConfig is the struct of the flomo-cli config.
type FlomoConfig struct {
	Api string `json:"api"`
}