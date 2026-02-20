package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var closeTabCmd = &cobra.Command{
	Use:   "close-tab",
	Short: "Close the specified tabs",
}

func init() {
	closeTabCmd.AddCommand(atCmd)
	carapace.Gen(closeTabCmd).Standalone()

	closeTabCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	closeTabCmd.Flags().Bool("ignore-no-match", false, "Do not return an error if no tabs are matched to be closed")
	closeTabCmd.Flags().StringP("match", "m", "", "The tab to match")
	closeTabCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	closeTabCmd.Flags().Bool("self", false, "Close the tab of the window this command is run in, rather than the active tab")

	carapace.Gen(closeTabCmd).FlagCompletion(carapace.ActionMap{})
}
