package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_alias_help_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Delete all aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_alias_help_clearCmd).Standalone()

	dotfiles_alias_helpCmd.AddCommand(dotfiles_alias_help_clearCmd)
}
