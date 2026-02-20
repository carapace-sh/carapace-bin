package cmd

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
	detachTabCmd.Flags().StringP("match", "m", "", "The tab to match")
	detachTabCmd.Flags().Bool("self", false, "Detach the tab this command is run in, rather than the active tab")
	detachTabCmd.Flags().StringP("target-tab", "t", "", "The tab to match")

	carapace.Gen(detachTabCmd).FlagCompletion(carapace.ActionMap{})
}
