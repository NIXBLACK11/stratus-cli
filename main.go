package main

import (
	"os"
	"stratus/commands"

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
		} else {
			color.Cyan(commands.Help())
		}
	} else if len(os.Args) == 3 {
		command := os.Args[1]
		if command == "project-name" {
			projectName := os.Args[2]
			color.Green("url list" + projectName)
		} else {
			color.Cyan(commands.Help())
		}
	} else {
		color.Cyan(commands.Help())
	}
}