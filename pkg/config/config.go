package config

import (
	"errors"
	"fmt"
	"os"
)

var (
	// Global that holds an instance of the Config struct, loaded using LoadConfig().
	container *configContainer
)

// Loads a configuration file from the current path, and interprets it as an instance of the typeof defaultConfig.
// If it doesn't find the file, it creates it.
func LoadOrCreateConfig(path string, defaultConfig interface{}) error {
	container = NewConfigContainer(defaultConfig)

	// Check that the config file doesn't exist and create it.
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err = container.Write(defaultConfig, path)
			if err != nil {
				return errors.New(fmt.Sprintf("cannot write config: %s", err.Error()))
			}
		}
	}

	_, err := container.Load(path)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot load config: %s", err.Error()))
	}

	return nil
}

// Returns the actual config instance.
func GetConfig() interface{} {
	return container.actualConfig
}
