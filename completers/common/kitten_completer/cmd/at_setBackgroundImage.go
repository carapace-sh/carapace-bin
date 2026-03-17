package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setBackgroundImageCmd = &cobra.Command{
	Use:   "set-background-image",
	Short: "Set the background image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setBackgroundImageCmd).Standalone()

	at_setBackgroundImageCmd.Flags().BoolP("all", "a", false, "By default, background image is only changed for the currently active OS window")
	at_setBackgroundImageCmd.Flags().BoolP("configured", "c", false, "Change the configured background image which is used for new OS windows")
	at_setBackgroundImageCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_setBackgroundImageCmd.Flags().String("layout", "configured", "How the image should be displayed")
	at_setBackgroundImageCmd.Flags().StringP("match", "m", "", "The window to match")
	at_setBackgroundImageCmd.Flags().Bool("no-response", false, "Don't wait for a response from kitty")
	atCmd.AddCommand(at_setBackgroundImageCmd)

}
