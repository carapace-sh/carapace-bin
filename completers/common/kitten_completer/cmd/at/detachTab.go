package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detachTabCmd = &cobra.Command{
	Use:   "detach-tab",
	Short: "Detach the specified tabs and place them in a different/new OS window",
}

func init() {
	detachTabCmd.AddCommand(atCmd)
	carapace.Gen(detachTabCmd).Standalone()

	detachTabCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(detachTabCmd).FlagCompletion(carapace.ActionMap{})
}