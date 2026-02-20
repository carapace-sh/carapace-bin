package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getTextCmd = &cobra.Command{
	Use:   "get-text",
	Short: "Get text from the specified window",
}

func init() {
	getTextCmd.AddCommand(atCmd)
	carapace.Gen(getTextCmd).Standalone()

	getTextCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(getTextCmd).FlagCompletion(carapace.ActionMap{})
}