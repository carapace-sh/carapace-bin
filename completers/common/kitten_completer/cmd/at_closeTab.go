package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_closeTabCmd = &cobra.Command{
	Use:   "close-tab",
	Short: "Close the specified tabs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_closeTabCmd).Standalone()

	at_closeTabCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_closeTabCmd.Flags().Bool("ignore-no-match", false, "Do not return an error if no tabs are matched to be closed")
	at_closeTabCmd.Flags().StringP("match", "m", "", "The tab to match")
	at_closeTabCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	at_closeTabCmd.Flags().Bool("self", false, "Close the tab of the window this command is run in, rather than the active tab")
	atCmd.AddCommand(at_closeTabCmd)

}
