package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var closeTabCmd = &cobra.Command{
	Use:   "closeTab",
	Short: "Close the specified tabs",
}

func init() {
	closeTabCmd.AddCommand(atCmd)
	carapace.Gen(closeTabCmd).Standalone()

	closeTabCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(closeTabCmd).FlagCompletion(carapace.ActionMap{})
}