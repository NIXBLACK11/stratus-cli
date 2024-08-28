package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"stratus-cli/types"
	"strconv"
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
	var tries int
	var alertTriggers []types.AlertTrigger
	var currentTrigger types.AlertTrigger

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "PROJECTNAME") {
			projectName = strings.TrimSpace(strings.TrimPrefix(line, "PROJECTNAME"))
		} else if strings.HasPrefix(line, "TRIES") {
			triesStr := strings.TrimSpace(strings.TrimPrefix(line, "TRIES"))
			tries, _ = strconv.Atoi(triesStr)
		}else if strings.HasPrefix(line, "SITENAME") {
			currentTrigger.SiteName = strings.TrimSpace(strings.TrimPrefix(line, "SITENAME"))
		} else if strings.HasPrefix(line, "SITEURL") {
			currentTrigger.SiteURL = strings.TrimSpace(strings.TrimPrefix(line, "SITEURL"))
		} else if strings.HasPrefix(line, "ALERTTYPE") {
			alertTypeStr := strings.TrimSpace(strings.TrimPrefix(line, "ALERTTYPE"))
			// Split the alert types by comma and trim spaces around each alert type
			alertTypes := strings.Split(alertTypeStr, ",")
			for i := range alertTypes {
				alertTypes[i] = strings.TrimSpace(alertTypes[i])
			}
			currentTrigger.AlertType = alertTypes
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
		Tries: tries,
		AlertTriggers: alertTriggers,
	}

	jsonData, err := json.MarshalIndent(projectDetails, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}

	return jsonData, nil
}
