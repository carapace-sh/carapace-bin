package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tab_getCmd = &cobra.Command{
	Use:   "get <tab_id>",
	Short: "Show a tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tab_getCmd).Standalone()

	tabCmd.AddCommand(tab_getCmd)

	carapace.Gen(tab_getCmd).PositionalCompletion(action.ActionTabs(tab_getCmd))
}
