package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var newWindowCmd = &cobra.Command{
	Use:   "new-window",
	Short: "Open new window",
}

func init() {
	newWindowCmd.AddCommand(atCmd)
	carapace.Gen(newWindowCmd).Standalone()

	newWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(newWindowCmd).FlagCompletion(carapace.ActionMap{})
}