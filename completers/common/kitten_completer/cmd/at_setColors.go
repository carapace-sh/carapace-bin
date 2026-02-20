package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setColorsCmd = &cobra.Command{
	Use:   "set-colors",
	Short: "Set terminal colors",
}

func init() {
	setColorsCmd.AddCommand(atCmd)
	carapace.Gen(setColorsCmd).Standalone()

	setColorsCmd.Flags().BoolP("all", "a", false, "By default, colors are only changed for the currently active window")
	setColorsCmd.Flags().BoolP("configured", "c", false, "Also change the configured colors")
	setColorsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	setColorsCmd.Flags().StringP("match", "m", "", "The window to match")
	setColorsCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	setColorsCmd.Flags().Bool("reset", false, "Restore all colors to the values they had at kitty startup")

	carapace.Gen(setColorsCmd).FlagCompletion(carapace.ActionMap{})
}