package utils

import (
	"os"
	"path/filepath"
)

func LoadEmail() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	emailFilePath := filepath.Join(homeDir, ".stratus_email")

	email, err := os.ReadFile(emailFilePath)
	if err != nil {
		return "", err
	}

	return string(email), nil
}
