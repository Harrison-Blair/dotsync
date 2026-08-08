package internal

import (
	"fmt"
	"os"
)

func List() ([]string, error) {
	var entries []string
	var err error

	dir, err := os.UserConfigDir()
	if err != nil {
		return entries, err
	}

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return entries, err
	}

	for _, entry := range dirEntries {
		fmt.Println(entry.Name())
		entries = append(entries, entry.Name())
	}

	return entries, err
}
