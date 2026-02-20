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

	setColorsCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setColorsCmd).FlagCompletion(carapace.ActionMap{})
}