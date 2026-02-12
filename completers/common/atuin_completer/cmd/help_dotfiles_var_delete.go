package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_dotfiles_var_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_dotfiles_var_deleteCmd).Standalone()

	help_dotfiles_varCmd.AddCommand(help_dotfiles_var_deleteCmd)
}
