package commands

import (
	"fmt"
	"stratus-cli/api"
	"stratus-cli/utils"

	"github.com/fatih/color"
)

func Signup() {
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
	
	success, err := api.Signup(email, password)
	if err!=nil {
		color.Red(err.Error())
		return
	}

	if success==true {
		color.Green("Successfully created user!!")
	}
}