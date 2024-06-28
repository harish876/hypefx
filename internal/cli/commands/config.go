package commands

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

type StaticFile struct {
	Package   string `json:"package"`
	Component string `json:"component"`
	FileName  string `json:"fileName"`
}

type HypeConfig struct {
	AppDir      string       `json:"appDir"`
	Module      string       `json:"module"`
	RoutesDir   string       `json:"routesDir"`
	Routing     bool         `json:"routing"`
	StaticFiles []StaticFile `json:"staticFiles"`
}

func SetConfig(newConfig HypeConfig) error {
	updatedJsonConfig, err := json.MarshalIndent(newConfig, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}
	if err := os.WriteFile(CONFIG_FILE_PATH, updatedJsonConfig, 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

func DeleteConfig(key string) error {
	var config map[string]interface{}
	if _, err := os.Stat(CONFIG_FILE_PATH); err == nil {
		file, err := os.ReadFile(CONFIG_FILE_PATH)
		if err != nil {
			return fmt.Errorf("failed to read file: %v", err)
		}

		if err := json.Unmarshal(file, &config); err != nil {
			return fmt.Errorf("failed to unmarshal JSON: %v", err)
		}
	} else {
		config = make(map[string]interface{})
	}

	delete(config, key)

	updatedConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	if err := os.WriteFile(CONFIG_FILE_PATH, updatedConfig, 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

func GetConfig() (HypeConfig, error) {
	var config HypeConfig
	if _, err := os.Stat(CONFIG_FILE_PATH); err != nil && os.IsNotExist(err) {
		slog.Error("GetAllConfig", "file stat error", err)
		if err := InitConfig(); err != nil {
			return config, err
		}
	}

	if _, err := os.Stat(CONFIG_FILE_PATH); err == nil {
		file, err := os.ReadFile(CONFIG_FILE_PATH)
		if err != nil {
			slog.Error("GetAllConfig", "config-file read", err)
			return config, err
		}

		if err := json.Unmarshal(file, &config); err != nil {
			return config, fmt.Errorf("failed to unmarshal JSON: %v", err)
		}
	}
	return config, nil
}

func InitConfig() error {
	f, err := os.Create(CONFIG_FILE_PATH)
	if err != nil {
		return err
	}
	defer f.Close()
	var config HypeConfig
	SetConfig(config)
	return nil
}
