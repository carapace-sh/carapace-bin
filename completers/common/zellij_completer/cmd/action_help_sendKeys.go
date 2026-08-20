package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_sendKeysCmd = &cobra.Command{
	Use:   "send-keys",
	Short: "Send one or more keys to the terminal (e.g., \"Ctrl a\", \"F1\", \"Alt Shift b\")",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_sendKeysCmd).Standalone()

	action_helpCmd.AddCommand(action_help_sendKeysCmd)
}
