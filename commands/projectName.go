package commands

import "stratus-cli/utils"

func ProjectDetails(projectName string) string {
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

	
}