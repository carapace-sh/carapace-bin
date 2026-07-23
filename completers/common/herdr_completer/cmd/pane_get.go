package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_getCmd = &cobra.Command{
	Use:   "get <pane_id>",
	Short: "Show a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_getCmd).Standalone()

	paneCmd.AddCommand(pane_getCmd)

	carapace.Gen(pane_getCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))
}
