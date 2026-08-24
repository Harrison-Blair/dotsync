package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFile = "config.json"

func Get(verbose bool) error {

	return nil
}

func Set() {

}

func readConfig(verbose bool) error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "dot", configFile)
	config, err := os.ReadFile(configPath)
	if errors.Is(err, os.ErrNotExist) {
		if verbose {
			fmt.Printf("no config file found at %q, creating one", configPath)
		}
		fmt.Printf("%q", config)
	}
	if err != nil {
		return fmt.Errorf("unable to read conifg file: %w", err)
	}

	if verbose {
		fmt.Printf("read config file: %q", configPath)
	}
	return nil
}
