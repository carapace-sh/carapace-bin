package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_dotfiles_aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Manage shell aliases with Atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_dotfiles_aliasCmd).Standalone()

	help_dotfilesCmd.AddCommand(help_dotfiles_aliasCmd)
}
