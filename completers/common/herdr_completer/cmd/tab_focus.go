package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
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

	carapace.Gen(tab_focusCmd).PositionalCompletion(herdr.ActionTabs(herdr.TabOpts{}))
}
