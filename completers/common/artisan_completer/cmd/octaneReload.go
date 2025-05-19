package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var octaneReloadCmd = &cobra.Command{
	Use:   "octane:reload [--server [SERVER]]",
	Short: "Reload the Octane workers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(octaneReloadCmd).Standalone()

	octaneReloadCmd.Flags().String("server", "", "The server that is running the application")
	rootCmd.AddCommand(octaneReloadCmd)
}
