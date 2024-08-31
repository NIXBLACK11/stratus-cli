package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"stratus-cli/constants"
	"stratus-cli/types"
)

func ProjectDetails(email string, token string, projectName string) (types.ProjectDetailsResponse, error) {
	var emptyDetails types.ProjectDetailsResponse

	if email == "" || token == "" || projectName == "" {
		return emptyDetails, errors.New("empty email, token, or project name")
	}

	URL := constants.URL
	if URL == "" {
		return emptyDetails, errors.New("error in loading env")
	}
	URL += email + "/" + projectName

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return emptyDetails, err
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return emptyDetails, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return emptyDetails, errors.New("bad response")
	}

	var details struct {
		ProjectDetails types.ProjectDetailsResponse `json:"projectdetails"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return emptyDetails, err
	}

	return details.ProjectDetails, nil
}
