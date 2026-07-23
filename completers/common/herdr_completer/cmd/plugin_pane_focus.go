package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
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

	carapace.Gen(plugin_pane_focusCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))
}
