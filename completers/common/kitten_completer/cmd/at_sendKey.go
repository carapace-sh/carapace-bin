package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sendKeyCmd = &cobra.Command{
	Use:   "send-key",
	Short: "Send arbitrary key presses to the specified windows",
}

func init() {
	sendKeyCmd.AddCommand(atCmd)
	carapace.Gen(sendKeyCmd).Standalone()

	sendKeyCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(sendKeyCmd).FlagCompletion(carapace.ActionMap{})
}