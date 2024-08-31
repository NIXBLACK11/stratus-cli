package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"stratus-cli/constants"
)

func Signup(email string, password string) (bool, error) {
	if email=="" || password=="" {
		return false, errors.New("Empty username or password")
	}

	URL := constants.URL
	if URL=="" {
		return false, errors.New("Error in loading env")
	}
	URL += "signup"

	requestData := map[string]string{
		"username": email,
		"password": password,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
    if err != nil {
        return false, err
    }

    req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK {
		return false, errors.New("Bad response")
	}
	
	return true, nil
}