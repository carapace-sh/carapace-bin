package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_togglePaneEmbedOrFloatingCmd = &cobra.Command{
	Use:   "toggle-pane-embed-or-floating",
	Short: "Embed focused pane if floating or float focused pane if embedded",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_togglePaneEmbedOrFloatingCmd).Standalone()

	help_actionCmd.AddCommand(help_action_togglePaneEmbedOrFloatingCmd)
}
