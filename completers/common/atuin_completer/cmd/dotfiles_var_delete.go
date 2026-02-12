package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var dotfiles_var_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_var_deleteCmd).Standalone()

	dotfiles_var_deleteCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_varCmd.AddCommand(dotfiles_var_deleteCmd)

	carapace.Gen(dotfiles_var_deleteCmd).PositionalCompletion(
		atuin.ActionVariables(),
	)
}
