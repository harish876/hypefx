package commands

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

type HypeConfig struct {
	AppDir    string `json:"app_dir"`
	Module    string `json:"module"`
	RoutesDir string `json:"routes_dir"`
	Routing   bool   `json:"routing"`
}

//TODO: add a static type validator later

func GetAllConfigV1() (*HypeConfig, error) {
	_, err := os.Stat(CONFIG_FILE_PATH)
	if err != nil && os.IsNotExist(err) {
		slog.Error("Config file does not exist")
		if err := InitConfig(); err != nil {
			return nil, err
		}
	}
	viper.SetConfigFile(CONFIG_FILE_PATH)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		slog.Debug("GetAllConfigV1", "Error reading hypeconfig file: %s\n", err)
		return nil, err
	}

	var config HypeConfig
	if err := viper.Unmarshal(&config); err != nil {
		slog.Debug("GetAllConfigV1", "failed to unmarshal JSON config %s\n", err)
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	slog.Info("GetAllConfigV1", "config struct", config)
	return &config, nil
}

func UpsertConfigV1(key string, value interface{}) error {
	viper.SetConfigFile(CONFIG_FILE_PATH)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	var existingVal interface{}
	if err := viper.UnmarshalKey(key, &existingVal); err != nil {
		return fmt.Errorf("error unmarshaling server config: %v", err)
	}

	viper.Set(key, value.(string))
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	slog.Info("UpsertConfigV1", "Config file updated successfully.", "Key", key, "Value", value)
	c, _ := GetAllConfigV1()
	slog.Debug("UpsertConfigV1----", "Config file.", c)
	return nil
}

func InitConfig() error {
	file, err := os.Create(CONFIG_FILE_PATH)
	if err != nil {
		return err
	}
	defer file.Close()
	viper.SetConfigName(CONFIG_FILE_PATH)
	viper.SetDefault("app_dir", "")
	viper.SetDefault("module", "")
	viper.SetDefault("routes_dir", "")
	viper.SetDefault("routing", true)

	if err := viper.WriteConfigAs(CONFIG_FILE_PATH); err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}
