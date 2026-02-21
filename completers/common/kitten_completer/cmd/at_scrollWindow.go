package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scrollWindowCmd = &cobra.Command{
	Use:   "scroll-window",
	Short: "Scroll the specified windows",
}

func init() {
	scrollWindowCmd.AddCommand(atCmd)
	carapace.Gen(scrollWindowCmd).Standalone()

	scrollWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	scrollWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	scrollWindowCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")

	carapace.Gen(scrollWindowCmd).FlagCompletion(carapace.ActionMap{})
}
