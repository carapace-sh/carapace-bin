package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sendTextCmd = &cobra.Command{
	Use:   "send-text",
	Short: "Send arbitrary text to specified windows",
}

func init() {
	sendTextCmd.AddCommand(atCmd)
	carapace.Gen(sendTextCmd).Standalone()

	sendTextCmd.Flags().Bool("all", false, "Match all windows")
	sendTextCmd.Flags().String("bracketed-paste", "disable", "When sending text to a window, wrap the text in bracketed paste escape codes")
	sendTextCmd.Flags().Bool("exclude-active", false, "Do not send text to the active window, even if it is one of the matched windows")
	sendTextCmd.Flags().String("from-file", "", "Path to a file whose contents you wish to send")
	sendTextCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	sendTextCmd.Flags().StringP("match", "m", "", "The window to match")
	sendTextCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	sendTextCmd.Flags().Bool("stdin", false, "Read the text to be sent from stdin")

	carapace.Gen(sendTextCmd).FlagCompletion(carapace.ActionMap{})
}
