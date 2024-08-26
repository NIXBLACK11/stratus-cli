package main

import (
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
			commands.Projects()
		} else if command == "login" {
			commands.Login()
		} else if command == "signup" {
			commands.Signup()
		} else {
			color.Cyan(commands.Help())
		}
	} else if len(os.Args) == 3 {
		command := os.Args[1]
		if command == "project" {
			projectName := os.Args[2]
			commands.ProjectDetails(projectName)
		} else if command == "delete" {
			projectName := os.Args[2]
			commands.ProjectDelete(projectName)
		} else {
			color.Cyan(commands.Help())
		}
	} else {
		color.Cyan(commands.Help())
	}
}