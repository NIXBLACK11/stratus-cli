package utils

import (
	"os"
	"path/filepath"
)

func StoreToken(token string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	tokenFilePath := filepath.Join(homeDir, ".stratus_token")
	file, err := os.OpenFile(tokenFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(token)
	if err != nil {
		return err
	}

	return nil
}
