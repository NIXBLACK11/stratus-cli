package commands

import (
	"fmt"
	"stratus-cli/api"

	"github.com/fatih/color"
)

func Login() {
	color.Cyan("Enter the email and password:\n")
	var email, password string
	color.Cyan("Email: ")
	fmt.Scanf("%s", &email)
	color.Cyan("Password: ")
	fmt.Scanf("%s", &password)
	
	success, err := api.Login(email, password)
	if err!=nil {
		color.Cyan(err.Error())
		return
	}

	if success==true {
		color.Green("Successfully logged in!!")
	}
}