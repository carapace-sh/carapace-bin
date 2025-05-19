package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var octaneStopCmd = &cobra.Command{
	Use:   "octane:stop [--server [SERVER]]",
	Short: "Stop the Octane server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(octaneStopCmd).Standalone()

	octaneStopCmd.Flags().String("server", "", "The server that is running the application")
	rootCmd.AddCommand(octaneStopCmd)
}
