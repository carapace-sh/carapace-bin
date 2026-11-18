package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_help_aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "List legacy shell aliases synced with Atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_help_aliasCmd).Standalone()

	dotfiles_helpCmd.AddCommand(dotfiles_help_aliasCmd)
}
