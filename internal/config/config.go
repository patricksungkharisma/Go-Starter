package config

import (
	"encoding/json"
	"log/slog"
	"os"

	errs "github.com/patricksungkharisma/go-starter/internal/error"
)

func InitConfig() (Config, error) {
	var (
		config Config
	)
	
	configJSON, err := os.ReadFile(developmentConfigFilePath)
	if err != nil {
		slog.Error(errs.ErrTagInitConfig,
			err,
			slog.Any("path", developmentConfigFilePath),
		)
		return config, err
	}

	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		slog.Error(errs.ErrTagInitConfig,
			err,
			slog.Any("config", config),
		)
		return config, err
	}

	return config, nil
}
