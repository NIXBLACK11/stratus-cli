package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"stratus-cli/constants"
)

type ProjectsResponse struct {
	Projects []string `json:"projects"`
}

func Projects(email string, token string) ([]string, error) {
	if email == "" || token == "" {
		return nil, errors.New("empty email or token")
	}

	URL := constants.URL
	if URL == "" {
		return nil, errors.New("error in loading env")
	}
	URL += email + "/projects"

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad response")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ProjectsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Projects, nil
}
