package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
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

	carapace.Gen(pane_getCmd).PositionalCompletion(action.ActionPanes(pane_getCmd))
}
