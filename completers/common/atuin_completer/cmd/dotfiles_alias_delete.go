package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var dotfiles_alias_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_alias_deleteCmd).Standalone()

	dotfiles_alias_deleteCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_aliasCmd.AddCommand(dotfiles_alias_deleteCmd)

	carapace.Gen(dotfiles_alias_deleteCmd).PositionalCompletion(
		atuin.ActionAliases(),
	)
}
