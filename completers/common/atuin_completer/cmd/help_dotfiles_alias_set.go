package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_dotfiles_alias_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_dotfiles_alias_setCmd).Standalone()

	help_dotfiles_aliasCmd.AddCommand(help_dotfiles_alias_setCmd)
}
