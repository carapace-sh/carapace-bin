package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selectWindowCmd = &cobra.Command{
	Use:   "select-window",
	Short: "Visually select a window in the specified tab",
}

func init() {
	selectWindowCmd.AddCommand(atCmd)
	carapace.Gen(selectWindowCmd).Standalone()

	selectWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(selectWindowCmd).FlagCompletion(carapace.ActionMap{})
}