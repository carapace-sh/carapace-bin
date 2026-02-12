package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_dotfiles_alias_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Delete all aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_dotfiles_alias_clearCmd).Standalone()

	help_dotfiles_aliasCmd.AddCommand(help_dotfiles_alias_clearCmd)
}
