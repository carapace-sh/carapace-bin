package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sailInstallCmd = &cobra.Command{
	Use:   "sail:install [--with [WITH]] [--devcontainer]",
	Short: "Install Laravel Sail's default Docker Compose file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sailInstallCmd).Standalone()

	sailInstallCmd.Flags().Bool("devcontainer", false, "Create a .devcontainer configuration directory")
	sailInstallCmd.Flags().String("with", "", "The services that should be included in the installation")
	rootCmd.AddCommand(sailInstallCmd)
}
