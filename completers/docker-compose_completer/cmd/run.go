package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a one-off command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	runCmd.Flags().StringS("e", "e", "", "Set an environment variable (can be used multiple times)")
	runCmd.Flags().BoolS("T", "T", false, "Disable pseudo-tty allocation. By default `docker-compose run`")
	runCmd.Flags().BoolP("detach", "d", false, "Detached mode: Run container in the background, print")
	runCmd.Flags().String("entrypoint", "", "Override the entrypoint of the image.")
	runCmd.Flags().StringP("label", "l", "", "Add or override a label (can be used multiple times)")
	runCmd.Flags().String("name", "", "Assign a name to the container")
	runCmd.Flags().Bool("no-deps", false, "Don't start linked services.")
	runCmd.Flags().StringP("publish", "p", "", "Publish a container's port(s) to the host")
	runCmd.Flags().Bool("rm", false, "Remove container after run. Ignored in detached mode.")
	runCmd.Flags().Bool("service-ports", false, "Run command with the service's ports enabled and mapped")
	runCmd.Flags().Bool("use-aliases", false, "Use the service's network aliases in the network(s) the")
	runCmd.Flags().StringP("user", "u", "", "Run as specified username or uid")
	runCmd.Flags().StringP("volume", "v", "", "Bind mount a volume (default [])")
	runCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalCompletion(
		ActionServices(),
	)
}
