package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tab_closeCmd = &cobra.Command{
	Use:   "close <tab_id>",
	Short: "Close a tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tab_closeCmd).Standalone()

	tabCmd.AddCommand(tab_closeCmd)

	carapace.Gen(tab_closeCmd).PositionalCompletion(action.ActionTabs(tab_closeCmd))
}
