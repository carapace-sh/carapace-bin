package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resizeWindowCmd = &cobra.Command{
	Use:   "resize-window",
	Short: "Resize the specified windows",
}

func init() {
	resizeWindowCmd.AddCommand(atCmd)
	carapace.Gen(resizeWindowCmd).Standalone()

	resizeWindowCmd.Flags().StringP("axis", "a", "horizontal", "The axis along which to resize")
	resizeWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	resizeWindowCmd.Flags().StringP("increment", "i", "2", "The number of cells to change the size by")
	resizeWindowCmd.Flags().StringP("match", "m", "", "The window to match")
	resizeWindowCmd.Flags().Bool("self", false, "Resize the window this command is run in, rather than the active window")

	carapace.Gen(resizeWindowCmd).FlagCompletion(carapace.ActionMap{})
}