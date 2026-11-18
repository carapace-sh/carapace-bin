package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_dotfiles_varCmd = &cobra.Command{
	Use:   "var",
	Short: "List legacy shell and environment variables synced with Atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_dotfiles_varCmd).Standalone()

	help_dotfilesCmd.AddCommand(help_dotfiles_varCmd)
}
