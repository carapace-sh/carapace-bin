package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tab_renameCmd = &cobra.Command{
	Use:   "rename <tab_id> <label>",
	Short: "Rename a tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tab_renameCmd).Standalone()

	tabCmd.AddCommand(tab_renameCmd)

	carapace.Gen(tab_renameCmd).PositionalCompletion(action.ActionTabs(tab_renameCmd))
}
