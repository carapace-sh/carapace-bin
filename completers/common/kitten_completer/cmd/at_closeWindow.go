package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_closeWindowCmd = &cobra.Command{
	Use:   "close-window",
	Short: "Close the specified windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_closeWindowCmd).Standalone()

	at_closeWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_closeWindowCmd.Flags().Bool("ignore-no-match", false, "Do not return an error if no windows are matched to be closed")
	at_closeWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	at_closeWindowCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	at_closeWindowCmd.Flags().Bool("self", false, "Close the window this command is run in, rather than the active window")
	atCmd.AddCommand(at_closeWindowCmd)

}
