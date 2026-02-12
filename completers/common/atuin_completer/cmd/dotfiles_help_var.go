package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_help_varCmd = &cobra.Command{
	Use:   "var",
	Short: "Manage shell and environment variables with Atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_help_varCmd).Standalone()

	dotfiles_helpCmd.AddCommand(dotfiles_help_varCmd)
}
