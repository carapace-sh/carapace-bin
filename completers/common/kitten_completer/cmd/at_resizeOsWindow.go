package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resizeOsWindowCmd = &cobra.Command{
	Use:   "resize-os-window",
	Short: "Resize/show/hide/etc. the specified OS Windows",
}

func init() {
	resizeOsWindowCmd.AddCommand(atCmd)
	carapace.Gen(resizeOsWindowCmd).Standalone()

	resizeOsWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(resizeOsWindowCmd).FlagCompletion(carapace.ActionMap{})
}