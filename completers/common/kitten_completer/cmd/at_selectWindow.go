package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selectWindowCmd = &cobra.Command{
	Use:   "select-window",
	Short: "Visually select a window in the specified tab",
}

func init() {
	selectWindowCmd.AddCommand(atCmd)
	carapace.Gen(selectWindowCmd).Standalone()

	selectWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	selectWindowCmd.Flags().Bool("exclude-active", false, "Exclude the currently active window from the list of windows to pick")
	selectWindowCmd.Flags().StringP("match", "m", "", "The tab to match")
	selectWindowCmd.Flags().Bool("reactivate-prev-tab", false, "When the selection is finished, the previously activated tab will be reactivated")
	selectWindowCmd.Flags().String("response-timeout", "60", "The time in seconds to wait for the user to select a window")
	selectWindowCmd.Flags().Bool("self", false, "Select window from the tab containing the window this command is run in")
	selectWindowCmd.Flags().String("title", "", "A title that will be displayed to the user to describe what this selection is for")

	carapace.Gen(selectWindowCmd).FlagCompletion(carapace.ActionMap{})
}