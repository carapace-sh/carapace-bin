package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var octaneStatusCmd = &cobra.Command{
	Use:   "octane:status [--server [SERVER]]",
	Short: "Get the current status of the Octane server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(octaneStatusCmd).Standalone()

	octaneStatusCmd.Flags().String("server", "", "The server that is running the application")
	rootCmd.AddCommand(octaneStatusCmd)
}
