package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detachWindowCmd = &cobra.Command{
	Use:   "detach-window",
	Short: "Detach the specified windows and place them in a different/new tab",
}

func init() {
	detachWindowCmd.AddCommand(atCmd)
	carapace.Gen(detachWindowCmd).Standalone()

	detachWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	detachWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	detachWindowCmd.Flags().Bool("self", false, "Detach the window this command is run in, rather than the active window")
	detachWindowCmd.Flags().Bool("stay-in-tab", false, "Keep the focus on a window in the currently focused tab after moving the specified windows")
	detachWindowCmd.Flags().StringP("target-tab", "t", "", "The tab to match")

	carapace.Gen(detachWindowCmd).FlagCompletion(carapace.ActionMap{})
}
