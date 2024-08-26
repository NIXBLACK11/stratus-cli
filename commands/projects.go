package commands

import (
	"stratus-cli/api"
	"stratus-cli/utils"

	"github.com/fatih/color"
)

func Projects() {
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

	projects, err := api.Projects(email, token)
	if err != nil {
		color.Red("Unable to make request at moment")
		return
	}

	if projects==nil {
		color.Cyan("No projects")
		return
	}

	for _, val := range projects {
		color.Cyan(val)
	}
}