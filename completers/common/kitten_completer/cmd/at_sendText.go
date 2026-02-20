package at

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

	sendTextCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(sendTextCmd).FlagCompletion(carapace.ActionMap{})
}