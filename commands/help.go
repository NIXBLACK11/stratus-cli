package commands

func Help() string {
	message := `Usage:
	help           Show available commands and options.
	projects       List all projects in your account.
	project-name   View and manage servers monitored in a specific project.`

	return message
}