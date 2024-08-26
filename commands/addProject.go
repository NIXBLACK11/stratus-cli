package commands

import (
	"stratus-cli/api"
	"stratus-cli/utils"

	"github.com/fatih/color"
)

func AddProject(projectFile string) {
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

	jsonData, err := utils.ExtractInfo(projectFile, email)
	if err != nil {
		color.Red("Invalid syntax in file"+projectFile)
		return
	}

	err = api.AddProject(email, token, jsonData)
	if err!=nil {
		color.Red(err.Error())
		return
	}

	color.Green("Successfully added project")
}