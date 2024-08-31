package api

import (
	"errors"
	"net/http"
	"stratus-cli/constants"
)

func ProjectDelete(email string, projectName string, token string) error {
	if email == "" || projectName == "" || token == "" {
		return errors.New("empty email, project name, or token")
	}

	URL := constants.URL
	if URL == "" {
		return errors.New("error in loading env")
	}
	URL += email + "/" + projectName

	req, err := http.NewRequest("DELETE", URL, nil)
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
		return errors.New("bad response")
	}

	return nil
}
