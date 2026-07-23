package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_closeCmd = &cobra.Command{
	Use:   "close <pane_id>",
	Short: "Close a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_closeCmd).Standalone()

	paneCmd.AddCommand(pane_closeCmd)

	carapace.Gen(pane_closeCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))
}
