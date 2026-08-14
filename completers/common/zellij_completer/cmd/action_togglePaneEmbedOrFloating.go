package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_togglePaneEmbedOrFloatingCmd = &cobra.Command{
	Use:   "toggle-pane-embed-or-floating",
	Short: "Embed focused pane if floating or float focused pane if embedded",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_togglePaneEmbedOrFloatingCmd).Standalone()

	action_togglePaneEmbedOrFloatingCmd.Flags().BoolP("help", "h", false, "Print help")
	action_togglePaneEmbedOrFloatingCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_togglePaneEmbedOrFloatingCmd)
}
