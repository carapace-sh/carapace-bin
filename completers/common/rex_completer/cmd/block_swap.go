package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var blockSwapCmd = &cobra.Command{
	Use:   "swap BLOCK OTHER",
	Short: "Swap the layout positions of two visible blocks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockSwapCmd).Standalone()

	common.AddWindowScopeFlags(blockSwapCmd)

	blockCmd.AddCommand(blockSwapCmd)

	carapace.Gen(blockSwapCmd).PositionalAnyCompletion(
		action.ActionBlocks(blockSwapCmd),
	)
}
