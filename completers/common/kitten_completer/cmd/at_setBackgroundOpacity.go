package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setBackgroundOpacityCmd = &cobra.Command{
	Use:   "set-background-opacity",
	Short: "Set the background opacity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setBackgroundOpacityCmd).Standalone()

	at_setBackgroundOpacityCmd.Flags().BoolP("all", "a", false, "By default, background opacity are only changed for the currently active OS window")
	at_setBackgroundOpacityCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_setBackgroundOpacityCmd.Flags().StringP("match", "m", "", "The window to match")
	at_setBackgroundOpacityCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	at_setBackgroundOpacityCmd.Flags().Bool("toggle", false, "When specified, the background opacity for the matching OS windows will be reset to default if it is currently equal to the specified value")
	atCmd.AddCommand(at_setBackgroundOpacityCmd)

}
