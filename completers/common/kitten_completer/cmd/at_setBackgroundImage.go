package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setBackgroundImageCmd = &cobra.Command{
	Use:   "set-background-image",
	Short: "Set the background image",
}

func init() {
	setBackgroundImageCmd.AddCommand(atCmd)
	carapace.Gen(setBackgroundImageCmd).Standalone()

	setBackgroundImageCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(setBackgroundImageCmd).FlagCompletion(carapace.ActionMap{})
}