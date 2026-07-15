package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tab_focusCmd = &cobra.Command{
	Use:   "focus <tab_id>",
	Short: "Focus a tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tab_focusCmd).Standalone()

	tabCmd.AddCommand(tab_focusCmd)

	carapace.Gen(tab_focusCmd).PositionalCompletion(action.ActionTabs(tab_focusCmd))
}
