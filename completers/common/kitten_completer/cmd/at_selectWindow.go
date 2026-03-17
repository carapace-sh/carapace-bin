package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_selectWindowCmd = &cobra.Command{
	Use:   "select-window",
	Short: "Visually select a window in the specified tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_selectWindowCmd).Standalone()

	at_selectWindowCmd.Flags().Bool("exclude-active", false, "Exclude the currently active window from the list of windows to pick")
	at_selectWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_selectWindowCmd.Flags().StringP("match", "m", "", "The tab to match")
	at_selectWindowCmd.Flags().Bool("reactivate-prev-tab", false, "When the selection is finished, the previously activated tab will be reactivated")
	at_selectWindowCmd.Flags().String("response-timeout", "60", "The time in seconds to wait for the user to select a window")
	at_selectWindowCmd.Flags().Bool("self", false, "Select window from the tab containing the window this command is run in")
	at_selectWindowCmd.Flags().String("title", "", "A title that will be displayed to the user to describe what this selection is for")
	atCmd.AddCommand(at_selectWindowCmd)

}
