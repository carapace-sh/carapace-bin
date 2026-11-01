package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_togglePaneEmbedOrFloatingCmd = &cobra.Command{
	Use:   "toggle-pane-embed-or-floating",
	Short: "Embed focused pane if floating or float focused pane if embedded",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_togglePaneEmbedOrFloatingCmd).Standalone()

	action_helpCmd.AddCommand(action_help_togglePaneEmbedOrFloatingCmd)
}
