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
	setWindowTitleCmd.Flags().StringP("match", "m", "", "The window to match")
	setWindowTitleCmd.Flags().Bool("temporary", false, "By default, the title will be permanently changed and programs running in the window will not be able to change it again")

	carapace.Gen(setWindowTitleCmd).FlagCompletion(carapace.ActionMap{})
}