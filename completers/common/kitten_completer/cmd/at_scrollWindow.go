package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scrollWindowCmd = &cobra.Command{
	Use:   "scroll-window",
	Short: "Scroll the specified windows",
}

func init() {
	scrollWindowCmd.AddCommand(atCmd)
	carapace.Gen(scrollWindowCmd).Standalone()

	scrollWindowCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(scrollWindowCmd).FlagCompletion(carapace.ActionMap{})
}