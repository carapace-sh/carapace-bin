package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var focusTabCmd = &cobra.Command{
	Use:   "focus-tab",
	Short: "Focus the specified tab",
}

func init() {
	focusTabCmd.AddCommand(atCmd)
	carapace.Gen(focusTabCmd).Standalone()

	focusTabCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(focusTabCmd).FlagCompletion(carapace.ActionMap{})
}