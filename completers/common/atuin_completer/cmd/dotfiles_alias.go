package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "List legacy shell aliases synced with Atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_aliasCmd).Standalone()

	dotfiles_aliasCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfilesCmd.AddCommand(dotfiles_aliasCmd)
}
