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

	sendKeyCmd.Flags().Bool("all", false, "Match all windows")
	sendKeyCmd.Flags().Bool("exclude-active", false, "Do not send text to the active window, even if it is one of the matched windows")
	sendKeyCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	sendKeyCmd.Flags().StringP("match", "m", "", "The window to match")
	sendKeyCmd.Flags().StringP("match-tab", "t", "", "The tab to match")

	carapace.Gen(sendKeyCmd).FlagCompletion(carapace.ActionMap{})
}
