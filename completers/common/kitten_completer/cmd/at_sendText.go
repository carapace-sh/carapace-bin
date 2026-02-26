package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_sendTextCmd = &cobra.Command{
	Use:   "send-text",
	Short: "Send arbitrary text to specified windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_sendTextCmd).Standalone()

	at_sendTextCmd.Flags().Bool("all", false, "Match all windows")
	at_sendTextCmd.Flags().String("bracketed-paste", "disable", "When sending text to a window, wrap the text in bracketed paste escape codes")
	at_sendTextCmd.Flags().Bool("exclude-active", false, "Do not send text to the active window, even if it is one of the matched windows")
	at_sendTextCmd.Flags().String("from-file", "", "Path to a file whose contents you wish to send")
	at_sendTextCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_sendTextCmd.Flags().StringP("match", "m", "", "The window to match")
	at_sendTextCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	at_sendTextCmd.Flags().Bool("stdin", false, "Read the text to be sent from stdin")
	atCmd.AddCommand(at_sendTextCmd)

}
