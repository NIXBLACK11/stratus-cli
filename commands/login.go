package commands

import (
	"fmt"
	"stratus-cli/api"
	"stratus-cli/utils"

	"github.com/fatih/color"
)

func Login() {
	color.Cyan("Enter the email and password:\n")
	var email, password string
	color.Cyan("Email: ")
	fmt.Scanf("%s", &email)

	validEmail, err := utils.EmailVerify(email)
	if validEmail==false {
		color.Red(err.Error())
		return
	}

	color.Cyan("Password: ")
	fmt.Scanf("%s", &password)
	
	token, err := api.Login(email, password)
	if err!=nil {
		color.Red(err.Error())
		return
	}

	err = utils.StoreToken(token)
	if err != nil {
		color.Red(err.Error())
		return
	}

	err = utils.StoreEmail(email)
	if err != nil {
		color.Red(err.Error())
		return
	}

	color.Green("Successfully logged in!!")
}