package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_detachTabCmd = &cobra.Command{
	Use:   "detach-tab",
	Short: "Detach the specified tabs and place them in a different/new OS window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_detachTabCmd).Standalone()

	at_detachTabCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_detachTabCmd.Flags().StringP("match", "m", "", "The tab to match")
	at_detachTabCmd.Flags().Bool("self", false, "Detach the tab this command is run in, rather than the active tab")
	at_detachTabCmd.Flags().StringP("target-tab", "t", "", "The tab to match")
	atCmd.AddCommand(at_detachTabCmd)

}
