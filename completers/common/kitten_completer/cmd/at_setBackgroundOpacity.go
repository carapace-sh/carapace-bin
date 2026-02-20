package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setBackgroundOpacityCmd = &cobra.Command{
	Use:   "set-background-opacity",
	Short: "Set the background opacity",
}

func init() {
	setBackgroundOpacityCmd.AddCommand(atCmd)
	carapace.Gen(setBackgroundOpacityCmd).Standalone()

	setBackgroundOpacityCmd.Flags().BoolP("all", "a", false, "By default, background opacity are only changed for the currently active OS window")
	setBackgroundOpacityCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	setBackgroundOpacityCmd.Flags().StringP("match", "m", "", "The window to match")
	setBackgroundOpacityCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	setBackgroundOpacityCmd.Flags().Bool("toggle", false, "When specified, the background opacity for the matching OS windows will be reset to default if it is currently equal to the specified value")

	carapace.Gen(setBackgroundOpacityCmd).FlagCompletion(carapace.ActionMap{})
}