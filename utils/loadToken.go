package utils

import (
	"os"
	"path/filepath"
)

func LoadToken() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	tokenFilePath := filepath.Join(homeDir, ".stratus_token")

	token, err := os.ReadFile(tokenFilePath)
	if err != nil {
		return "", err
	}

	return string(token), nil
}
