package at

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

	setBackgroundOpacityCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setBackgroundOpacityCmd).FlagCompletion(carapace.ActionMap{})
}