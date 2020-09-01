package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Get help on a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"build", "Build or rebuild services",
			"config", "Validate and view the Compose file",
			"create", "Create services",
			"down", "Stop and remove containers, networks, images, and volumes",
			"events", "Receive real time events from containers",
			"exec", "Execute a command in a running container",
			"help", "Get help on a command",
			"images", "List images",
			"kill", "Kill containers",
			"logs", "View output from containers",
			"pause", "Pause services",
			"port", "Print the public port for a port binding",
			"ps", "List containers",
			"pull", "Pull service images",
			"push", "Push service images",
			"restart", "Restart services",
			"rm", "Remove stopped containers",
			"run", "Run a one-off command",
			"scale", "Set number of containers for a service",
			"start", "Start services",
			"stop", "Stop services",
			"top", "Display the running processes",
			"unpause", "Unpause services",
			"up", "Create and start containers",
			"version", "Show the Docker-Compose version information",
		),
	)
}
