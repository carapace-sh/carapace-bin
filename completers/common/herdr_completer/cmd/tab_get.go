package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
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

	carapace.Gen(tab_getCmd).PositionalCompletion(herdr.ActionTabs(herdr.TabOpts{}))
}
