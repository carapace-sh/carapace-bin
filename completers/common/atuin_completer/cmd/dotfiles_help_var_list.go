package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_help_var_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_help_var_listCmd).Standalone()

	dotfiles_help_varCmd.AddCommand(dotfiles_help_var_listCmd)
}
