package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detachWindowCmd = &cobra.Command{
	Use:   "detach-window",
	Short: "Detach the specified windows and place them in a different/new tab",
}

func init() {
	detachWindowCmd.AddCommand(atCmd)
	carapace.Gen(detachWindowCmd).Standalone()

	detachWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(detachWindowCmd).FlagCompletion(carapace.ActionMap{})
}