package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_alias_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_alias_help_listCmd).Standalone()

	dotfiles_alias_helpCmd.AddCommand(dotfiles_alias_help_listCmd)
}
