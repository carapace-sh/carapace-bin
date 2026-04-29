package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate [shell]",
	Short: "Initialize mise in the current shell session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(activateCmd).Standalone()
	activateCmd.Flags().Bool("shims", false, "Use shims instead of shell functions")
	rootCmd.AddCommand(activateCmd)

	carapace.Gen(activateCmd).PositionalCompletion(
		carapace.ActionValues("bash", "elvish", "fish", "nu", "xonsh", "zsh"),
	)
}
