package list

import (
	"fmt"
	"os"
	"strings"
)

func List(verbose bool) ([]os.DirEntry, error) {
	var entries []os.DirEntry

	candidates, err := getCandidates(verbose)
	if err != nil {
		return entries, fmt.Errorf("failed to get candidates: %w", err)
	}
	entries = append(entries, candidates...)

	return entries, nil
}

// TODO:
//   - Add custom config support
//   - Allow home dir to be ommited (both by config, and on error)
func getCandidates(verbose bool) ([]os.DirEntry, error) {
	var candidates []os.DirEntry

	// Config Directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return candidates, fmt.Errorf("failed to get config directory: %w", err)
	}
	if verbose {
		fmt.Printf("found config directory: %q\n", configDir)
	}

	configEntires, err := os.ReadDir(configDir)
	if err != nil {
		return candidates, fmt.Errorf("failed to read %q: %w", configDir, err)
	}
	if verbose {
		fmt.Printf("found %d config entries (files & directories)\n", len(configEntires))
	}
	candidates = append(candidates, configEntires...)

	// Home Directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return candidates, fmt.Errorf("failed to get home directory: %w", err)
	}
	if verbose {
		fmt.Printf("found home directory: %q\n", homeDir)
	}

	homeEntries, err := os.ReadDir(homeDir)
	if err != nil {
		return candidates, fmt.Errorf("failed to read %q: %w", homeDir, err)
	}
	if verbose {
		fmt.Printf("config directories found in %q:\n", homeDir)
	}
	for _, homeEntry := range homeEntries {
		if strings.HasPrefix(homeEntry.Name(), ".") {
			if verbose {
				fmt.Printf("  %q\n", homeEntry.Name())
			}
			candidates = append(candidates, homeEntry)
		}
	}

	if verbose {
		fmt.Printf("found %d candidates total\n", len(candidates))
	}

	return candidates, nil
}
