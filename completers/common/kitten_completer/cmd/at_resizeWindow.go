package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_resizeWindowCmd = &cobra.Command{
	Use:   "resize-window",
	Short: "Resize the specified windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_resizeWindowCmd).Standalone()

	at_resizeWindowCmd.Flags().StringP("axis", "a", "horizontal", "The axis along which to resize")
	at_resizeWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_resizeWindowCmd.Flags().StringP("increment", "i", "2", "The number of cells to change the size by")
	at_resizeWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	at_resizeWindowCmd.Flags().Bool("self", false, "Resize the window this command is run in, rather than the active window")
	atCmd.AddCommand(at_resizeWindowCmd)

}
