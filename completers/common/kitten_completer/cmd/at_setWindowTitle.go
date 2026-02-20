package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setWindowTitleCmd = &cobra.Command{
	Use:   "set-window-title",
	Short: "Set the window title",
}

func init() {
	setWindowTitleCmd.AddCommand(atCmd)
	carapace.Gen(setWindowTitleCmd).Standalone()

	setWindowTitleCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setWindowTitleCmd).FlagCompletion(carapace.ActionMap{})
}