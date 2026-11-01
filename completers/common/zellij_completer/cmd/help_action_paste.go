package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "Paste text to the terminal (using bracketed paste mode)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_pasteCmd).Standalone()

	help_actionCmd.AddCommand(help_action_pasteCmd)
}
