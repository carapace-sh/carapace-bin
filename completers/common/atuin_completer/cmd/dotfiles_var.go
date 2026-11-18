package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_varCmd = &cobra.Command{
	Use:   "var",
	Short: "List legacy shell and environment variables synced with Atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_varCmd).Standalone()

	dotfiles_varCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfilesCmd.AddCommand(dotfiles_varCmd)
}
