package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_pane_focusCmd = &cobra.Command{
	Use:   "focus <pane_id>",
	Short: "Focus a plugin pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_pane_focusCmd).Standalone()

	plugin_paneCmd.AddCommand(plugin_pane_focusCmd)

	carapace.Gen(plugin_pane_focusCmd).PositionalCompletion(action.ActionPanes(plugin_pane_focusCmd))
}
