package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "Paste text to the terminal (using bracketed paste mode)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_pasteCmd).Standalone()

	action_helpCmd.AddCommand(action_help_pasteCmd)
}
