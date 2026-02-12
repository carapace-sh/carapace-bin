package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_help_alias_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_help_alias_deleteCmd).Standalone()

	dotfiles_help_aliasCmd.AddCommand(dotfiles_help_alias_deleteCmd)
}
