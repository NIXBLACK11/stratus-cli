package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type AlertTrigger struct {
	SiteName  string `json:"sitename"`
	SiteURL   string `json:"siteurl"`
	AlertType string `json:"alerttype"`
}

type ProjectDetailsResponse struct {
	Username     string        `json:"username"`
	ProjectName  string        `json:"projectname"`
	AlertTriggers []AlertTrigger `json:"AlertTriggers"`
}

func ProjectDetails(email string, token string, projectName string) (*ProjectDetailsResponse, error) {
	if email == "" || token == "" || projectName == "" {
		return nil, errors.New("empty email, token, or project name")
	}

	URL := os.Getenv("URL")
	if URL == "" {
		return nil, errors.New("error in loading env")
	}
	URL += email + "/" + projectName

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

	// Read the response body and unmarshal it into the ProjectDetailsResponse struct
	var details struct {
		ProjectDetails ProjectDetailsResponse `json:"projectdetails"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, err
	}

	return &details.ProjectDetails, nil
}
