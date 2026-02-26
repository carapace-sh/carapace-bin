package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_detachWindowCmd = &cobra.Command{
	Use:   "detach-window",
	Short: "Detach the specified windows and place them in a different/new tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_detachWindowCmd).Standalone()

	at_detachWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_detachWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	at_detachWindowCmd.Flags().Bool("self", false, "Detach the window this command is run in, rather than the active window")
	at_detachWindowCmd.Flags().Bool("stay-in-tab", false, "Keep the focus on a window in the currently focused tab after moving the specified windows")
	at_detachWindowCmd.Flags().StringP("target-tab", "t", "", "The tab to match")
	atCmd.AddCommand(at_detachWindowCmd)

}
