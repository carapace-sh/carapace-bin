package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setWindowLogoCmd = &cobra.Command{
	Use:   "set-window-logo",
	Short: "Set the window logo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setWindowLogoCmd).Standalone()

	at_setWindowLogoCmd.Flags().String("alpha", "-1", "The amount the window logo should be faded into the background")
	at_setWindowLogoCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_setWindowLogoCmd.Flags().StringP("match", "m", "", "The window to match")
	at_setWindowLogoCmd.Flags().Bool("no-response", false, "Don't wait for a response from kitty")
	at_setWindowLogoCmd.Flags().String("position", "", "The position for the window logo")
	at_setWindowLogoCmd.Flags().Bool("self", false, "Act on the window this command is run in, rather than the active window")
	atCmd.AddCommand(at_setWindowLogoCmd)

}
