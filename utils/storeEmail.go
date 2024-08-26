package utils

import (
	"os"
	"path/filepath"
)

func StoreEmail(email string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	emailFilePath := filepath.Join(homeDir, ".stratus_email")
	file, err := os.OpenFile(emailFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(email)
	if err != nil {
		return err
	}

	return nil
}
