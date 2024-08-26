package commands

import (
	"stratus-cli/api"
	"stratus-cli/utils"

	"github.com/fatih/color"
)
func ProjectDelete(projectName string) {
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

	err = api.ProjectDelete(email, projectName, token)
	if err != nil {
		color.Red("Unable to make request at moment")
		return
	}

	color.Green("Project deleted successfully")
}