package commands

import (
	"fmt"
	"stratus-cli/api"
	"stratus-cli/utils"

	"github.com/fatih/color"
)

func ProjectDetails(projectName string) {
	email, err := utils.LoadEmail()
	if err != nil {
		color.Red("User not signed in!!")
		return
	}

	token, err := utils.LoadToken()
	if err != nil {
		color.Red("User not signed in!!")
		return
	}

	details, err := api.ProjectDetails(email, token, projectName)
	if err != nil {
		color.Red(err.Error())
		return
	}

	color.Green("Project Details:")
	color.Green(fmt.Sprintf("Username: %s\n", details.Username))
	color.Green(fmt.Sprintf("Project Name: %s\n", details.ProjectName))
	color.Green("Alert Triggers:")
	for i, trigger := range details.AlertTriggers {
		color.Green(fmt.Sprintf("  %d. Site Name: %s\n", i+1, trigger.SiteName))
		color.Green(fmt.Sprintf("     Site URL: %s\n", trigger.SiteURL))
		color.Green(fmt.Sprintf("     Alert Type: %s\n", trigger.AlertType))
	}
}