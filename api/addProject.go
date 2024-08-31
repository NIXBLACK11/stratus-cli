package api

import (
	"bytes"
	"errors"
	"net/http"
	"stratus-cli/constants"
)

func AddProject(email string, token string, jsonData []byte) error {
	if email == "" || token == "" {
		return errors.New("empty email or token")
	}

	URL := constants.URL
	if URL == "" {
		return errors.New("error in loading env")
	}
	URL += email + "/addProject"

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("bad response: " + resp.Status)
	}

	return nil
}
