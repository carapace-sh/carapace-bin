package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var swapCmd = &cobra.Command{
	Use:   "swap BLOCK OTHER",
	Short: "Swap the layout positions of two visible blocks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swapCmd).Standalone()

	common.AddWindowScopeFlags(swapCmd)

	rootCmd.AddCommand(swapCmd)

	carapace.Gen(swapCmd).PositionalAnyCompletion(
		action.ActionBlocks(swapCmd),
	)
}
