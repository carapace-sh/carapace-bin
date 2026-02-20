package cmd

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

	getTextCmd.Flags().Bool("add-cursor", false, "Add ANSI escape codes specifying the cursor position and style to the end of the text")
	getTextCmd.Flags().Bool("add-wrap-markers", false, "Add carriage returns at every line wrap location")
	getTextCmd.Flags().Bool("ansi", false, "Include the ANSI formatting escape codes for colors, bold, italic, etc.")
	getTextCmd.Flags().Bool("clear-selection", false, "Clear the selection in the matched window, if any")
	getTextCmd.Flags().String("extent", "screen", "What text to get")
	getTextCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	getTextCmd.Flags().StringP("match", "m", "", "The window to match")
	getTextCmd.Flags().Bool("self", false, "Get text from the window this command is run in, rather than the active window")

	carapace.Gen(getTextCmd).FlagCompletion(carapace.ActionMap{
		"extent": carapace.ActionValues("screen", "all", "first_cmd_output_on_screen", "last_cmd_output", "last_non_empty_output", "last_visited_cmd_output", "selection"),
	})
}