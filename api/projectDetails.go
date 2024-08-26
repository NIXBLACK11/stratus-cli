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

func ProjectDetails(email string, token string, projectName string) (ProjectDetailsResponse, error) {
	var emptyDetails ProjectDetailsResponse

	if email == "" || token == "" || projectName == "" {
		return emptyDetails, errors.New("empty email, token, or project name")
	}

	URL := os.Getenv("URL")
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
		ProjectDetails ProjectDetailsResponse `json:"projectdetails"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return emptyDetails, err
	}

	return details.ProjectDetails, nil
}
