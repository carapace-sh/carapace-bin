package at

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

	resizeWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(resizeWindowCmd).FlagCompletion(carapace.ActionMap{})
}