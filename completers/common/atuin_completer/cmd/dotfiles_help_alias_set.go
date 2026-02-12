package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_help_alias_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_help_alias_setCmd).Standalone()

	dotfiles_help_aliasCmd.AddCommand(dotfiles_help_alias_setCmd)
}
