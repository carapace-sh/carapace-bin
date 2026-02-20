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

	setBackgroundImageCmd.Flags().BoolP("all", "a", false, "By default, background image is only changed for the currently active OS window")
	setBackgroundImageCmd.Flags().BoolP("configured", "c", false, "Change the configured background image which is used for new OS windows")
	setBackgroundImageCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	setBackgroundImageCmd.Flags().String("layout", "configured", "How the image should be displayed")
	setBackgroundImageCmd.Flags().StringP("match", "m", "", "The window to match")
	setBackgroundImageCmd.Flags().Bool("no-response", false, "Don't wait for a response from kitty")

	carapace.Gen(setBackgroundImageCmd).FlagCompletion(carapace.ActionMap{})
}