package config

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

func InitConfig() (Config, error) {
	config := Config{}
	configJSON, err := os.ReadFile(configFilePath)
	if err != nil {
		return config, errors.Wrap(err, "[InitConfig]")
	}

	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		return config, errors.Wrap(err, "[InitConfig]")
	}

	return config, nil
}
