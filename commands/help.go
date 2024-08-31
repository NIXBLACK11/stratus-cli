package commands

func Help() string {
	message := `Usage:
	    help                         Show available commands and options.
        login                        Login to your account
        signup                       Create a new account
        projects                     List all projects in your account.
        project <project_name>       View and manage urls monitored in a specific project.
        delete <project_name>        Delete the prject specified
        add <config.st>              Add details to an existing or create a new project`
	return message
}