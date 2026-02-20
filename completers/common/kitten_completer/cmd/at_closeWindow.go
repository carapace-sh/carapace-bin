package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var closeWindowCmd = &cobra.Command{
	Use:   "close-window",
	Short: "Close the specified windows",
}

func init() {
	closeWindowCmd.AddCommand(atCmd)
	carapace.Gen(closeWindowCmd).Standalone()

	closeWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(closeWindowCmd).FlagCompletion(carapace.ActionMap{})
}