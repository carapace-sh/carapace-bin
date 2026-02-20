package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var closeWindowCmd = &cobra.Command{
	Use:   "close-window",
	Short: "Close the specified windows",
}

func init() {
	closeWindowCmd.AddCommand(atCmd)
	carapace.Gen(closeWindowCmd).Standalone()

	closeWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	closeWindowCmd.Flags().Bool("ignore-no-match", false, "Do not return an error if no windows are matched to be closed")
	closeWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	closeWindowCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	closeWindowCmd.Flags().Bool("self", false, "Close the window this command is run in, rather than the active window")

	carapace.Gen(closeWindowCmd).FlagCompletion(carapace.ActionMap{})
}