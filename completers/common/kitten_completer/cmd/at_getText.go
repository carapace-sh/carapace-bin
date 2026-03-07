package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_getTextCmd = &cobra.Command{
	Use:   "get-text",
	Short: "Get text from the specified window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_getTextCmd).Standalone()

	at_getTextCmd.Flags().Bool("add-cursor", false, "Add ANSI escape codes specifying the cursor position and style to the end of the text")
	at_getTextCmd.Flags().Bool("add-wrap-markers", false, "Add carriage returns at every line wrap location")
	at_getTextCmd.Flags().Bool("ansi", false, "Include the ANSI formatting escape codes for colors, bold, italic, etc.")
	at_getTextCmd.Flags().Bool("clear-selection", false, "Clear the selection in the matched window, if any")
	at_getTextCmd.Flags().String("extent", "screen", "What text to get")
	at_getTextCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_getTextCmd.Flags().StringP("match", "m", "", "The window to match")
	at_getTextCmd.Flags().Bool("self", false, "Get text from the window this command is run in, rather than the active window")
	atCmd.AddCommand(at_getTextCmd)

	carapace.Gen(at_getTextCmd).FlagCompletion(carapace.ActionMap{
		"extent": carapace.ActionValues("screen", "all", "first_cmd_output_on_screen", "last_cmd_output", "last_non_empty_output", "last_visited_cmd_output", "selection"),
	})
}
