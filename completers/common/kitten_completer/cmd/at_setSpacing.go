package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setSpacingCmd = &cobra.Command{
	Use:   "set-spacing",
	Short: "Set window paddings and margins",
}

func init() {
	setSpacingCmd.AddCommand(atCmd)
	carapace.Gen(setSpacingCmd).Standalone()

	setSpacingCmd.Flags().BoolP("all", "a", false, "By default, settings are only changed for the currently active window")
	setSpacingCmd.Flags().BoolP("configured", "c", false, "Also change the configured paddings and margins")
	setSpacingCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	setSpacingCmd.Flags().StringP("match", "m", "", "The window to match")
	setSpacingCmd.Flags().StringP("match-tab", "t", "", "The tab to match")

	carapace.Gen(setSpacingCmd).FlagCompletion(carapace.ActionMap{})
}
