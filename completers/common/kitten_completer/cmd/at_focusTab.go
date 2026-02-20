package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var focusTabCmd = &cobra.Command{
	Use:   "focus-tab",
	Short: "Focus the specified tab",
}

func init() {
	focusTabCmd.AddCommand(atCmd)
	carapace.Gen(focusTabCmd).Standalone()

	focusTabCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	focusTabCmd.Flags().StringP("match", "m", "", "The tab to match")
	focusTabCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")

	carapace.Gen(focusTabCmd).FlagCompletion(carapace.ActionMap{})
}