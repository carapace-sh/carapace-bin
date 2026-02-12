package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_var_help_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_var_help_deleteCmd).Standalone()

	dotfiles_var_helpCmd.AddCommand(dotfiles_var_help_deleteCmd)
}
