package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_dotfilesCmd = &cobra.Command{
	Use:   "dotfiles",
	Short: "Manage your dotfiles with Atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_dotfilesCmd).Standalone()

	helpCmd.AddCommand(help_dotfilesCmd)
}
