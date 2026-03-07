package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_sendKeyCmd = &cobra.Command{
	Use:   "send-key",
	Short: "Send arbitrary key presses to the specified windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_sendKeyCmd).Standalone()

	at_sendKeyCmd.Flags().Bool("all", false, "Match all windows")
	at_sendKeyCmd.Flags().Bool("exclude-active", false, "Do not send text to the active window, even if it is one of the matched windows")
	at_sendKeyCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_sendKeyCmd.Flags().StringP("match", "m", "", "The window to match")
	at_sendKeyCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	atCmd.AddCommand(at_sendKeyCmd)

}
