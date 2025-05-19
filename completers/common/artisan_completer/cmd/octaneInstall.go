package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var octaneInstallCmd = &cobra.Command{
	Use:   "octane:install [--server [SERVER]]",
	Short: "Install the Octane components and resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(octaneInstallCmd).Standalone()

	octaneInstallCmd.Flags().String("server", "", "The server that should be used to serve the application")
	rootCmd.AddCommand(octaneInstallCmd)
}
