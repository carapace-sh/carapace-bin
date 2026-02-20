package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setWindowLogoCmd = &cobra.Command{
	Use:   "set-window-logo",
	Short: "Set the window logo",
}

func init() {
	setWindowLogoCmd.AddCommand(atCmd)
	carapace.Gen(setWindowLogoCmd).Standalone()

	setWindowLogoCmd.Flags().String("alpha", "-1", "The amount the window logo should be faded into the background")
	setWindowLogoCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	setWindowLogoCmd.Flags().StringP("match", "m", "", "The window to match")
	setWindowLogoCmd.Flags().Bool("no-response", false, "Don't wait for a response from kitty")
	setWindowLogoCmd.Flags().String("position", "", "The position for the window logo")
	setWindowLogoCmd.Flags().Bool("self", false, "Act on the window this command is run in, rather than the active window")

	carapace.Gen(setWindowLogoCmd).FlagCompletion(carapace.ActionMap{})
}
