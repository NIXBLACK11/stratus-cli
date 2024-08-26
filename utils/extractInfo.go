package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"stratus-cli/types"
	"os"
	"strings"
)

func ExtractInfo(projectFile string, username string) ( []byte, error ) {
	file, err := os.Open(projectFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var projectName string
	var alertTriggers []types.AlertTrigger
	var currentTrigger types.AlertTrigger

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "PROJECTNAME") {
			projectName = strings.TrimSpace(strings.TrimPrefix(line, "PROJECTNAME"))
		} else if strings.HasPrefix(line, "SITENAME") {
			currentTrigger.SiteName = strings.TrimSpace(strings.TrimPrefix(line, "SITENAME"))
		} else if strings.HasPrefix(line, "SITEURL") {
			currentTrigger.SiteURL = strings.TrimSpace(strings.TrimPrefix(line, "SITEURL"))
		} else if strings.HasPrefix(line, "ALERTTYPE") {
			currentTrigger.AlertType = strings.TrimSpace(strings.TrimPrefix(line, "ALERTTYPE"))
			alertTriggers = append(alertTriggers, currentTrigger)
			currentTrigger = types.AlertTrigger{}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	projectDetails := types.ProjectDetailsResponse{
		Username:     username,
		ProjectName:  projectName,
		AlertTriggers: alertTriggers,
	}

	jsonData, err := json.MarshalIndent(projectDetails, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}

	return jsonData, nil
}
