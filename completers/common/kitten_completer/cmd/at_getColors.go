package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getColorsCmd = &cobra.Command{
	Use:   "get-colors",
	Short: "Get terminal colors",
}

func init() {
	getColorsCmd.AddCommand(atCmd)
	carapace.Gen(getColorsCmd).Standalone()

	getColorsCmd.Flags().BoolP("configured", "c", false, "Instead of outputting the colors for the specified window, output the currently configured colors")
	getColorsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	getColorsCmd.Flags().StringP("match", "m", "", "The window to match")

	carapace.Gen(getColorsCmd).FlagCompletion(carapace.ActionMap{})
}
