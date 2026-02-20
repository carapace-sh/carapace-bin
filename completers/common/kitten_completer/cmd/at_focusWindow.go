package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var focusWindowCmd = &cobra.Command{
	Use:   "focus-window",
	Short: "Focus the specified window",
}

func init() {
	focusWindowCmd.AddCommand(atCmd)
	carapace.Gen(focusWindowCmd).Standalone()

	focusWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	focusWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	focusWindowCmd.Flags().Bool("no-response", false, "Don't wait for a response from kitty")

	carapace.Gen(focusWindowCmd).FlagCompletion(carapace.ActionMap{})
}