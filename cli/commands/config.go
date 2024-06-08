package commands

import (
	"encoding/json"
	"fmt"
	"os"
)

func UpsertConfig(key string, value interface{}) error {
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

	config[key] = value

	updatedConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	if err := os.WriteFile(CONFIG_FILE_PATH, updatedConfig, 0644); err != nil {
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

	if _, ok := config[key]; ok {
		delete(config, key)
	}

	updatedConfig, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	if err := os.WriteFile(CONFIG_FILE_PATH, updatedConfig, 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

func GetConfig(key string) (interface{}, error) {
	var config map[string]interface{}
	if _, err := os.Stat(CONFIG_FILE_PATH); err == nil {
		file, err := os.ReadFile(CONFIG_FILE_PATH)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %v", err)
		}

		if err := json.Unmarshal(file, &config); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
		}
	} else {
		config = make(map[string]interface{})
	}

	if value, ok := config[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("could not find config with key %s", key)
}
