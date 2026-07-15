package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_pane_closeCmd = &cobra.Command{
	Use:   "close <pane_id>",
	Short: "Close a plugin pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_pane_closeCmd).Standalone()

	plugin_paneCmd.AddCommand(plugin_pane_closeCmd)

	carapace.Gen(plugin_pane_closeCmd).PositionalCompletion(action.ActionPanes(plugin_pane_closeCmd))
}
