package commands

func Help() string {
	message := `Usage:
	help           Show available commands and options.
	login          Login to your account
	signup         Create a new account
	projects       List all projects in your account.
	project        View and manage urls monitored in a specific project.
	delete         Delete the prject specified`

	return message
}