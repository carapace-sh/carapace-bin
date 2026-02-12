package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var dotfiles_var_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_var_setCmd).Standalone()

	dotfiles_var_setCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_var_setCmd.Flags().BoolP("no-export", "n", false, "")
	dotfiles_varCmd.AddCommand(dotfiles_var_setCmd)

	carapace.Gen(dotfiles_var_setCmd).PositionalCompletion(
		atuin.ActionVariables(),
	)
}
