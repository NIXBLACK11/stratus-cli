package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"stratus-cli/constants"
)

func Login(email string, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("empty username or password")
	}

	URL := constants.URL
	if URL == "" {
		return "", errors.New("error in loading env")
	}
	URL += "login"

	requestData := map[string]string{
		"username": email,
		"password": password,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("bad response")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseMap map[string]string
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		return "", err
	}

	token, ok := responseMap["token"]
	if !ok {
		return "", errors.New("Api does not return token")
	}

	return token, nil
}
