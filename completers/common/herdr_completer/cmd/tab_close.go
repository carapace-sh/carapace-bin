package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
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

	carapace.Gen(tab_closeCmd).PositionalCompletion(herdr.ActionTabs(herdr.TabOpts{}))
}
