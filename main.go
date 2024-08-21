package main

import (
	"fmt"
	"os"
	"stratus-cli/commands"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		color.Cyan(commands.Help())
	} else if len(os.Args) == 2 {
		command := os.Args[1]
		if command == "help" {
			color.Cyan(commands.Help())
		} else if command == "projects" {
			color.Blue("projects list")
		} else if command == "login" {
			color.Cyan("Enter the email and password:\n")
			var email, password string
			fmt.Printf("Email: ")
			fmt.Scanf("%s", &email)
			fmt.Printf("Password: ")
			fmt.Scanf("%s", &password)
			success := commands.Login(email, password)
			if success==false {
				fmt.Println("Invalid email or password")
			}
			fmt.Println("Logged in")
		} else if command == "signup" {
			color.Cyan("Enter the email and password:\n")
		} else {
			color.Cyan(commands.Help())
		}
	} else if len(os.Args) == 3 {
		command := os.Args[1]
		if command == "project" {
			projectName := os.Args[2]
			color.Green("url list" + projectName)
		} else if command == "delete" {
			projectName := os.Args[2]
			color.Green("deleted" + projectName)
		} else {
			color.Cyan(commands.Help())
		}
	} else {
		color.Cyan(commands.Help())
	}
}