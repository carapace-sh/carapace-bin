package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var blockCloseCmd = &cobra.Command{
	Use:   "close BLOCK...",
	Short: "Close blocks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockCloseCmd).Standalone()

	common.AddWindowScopeFlags(blockCloseCmd)

	blockCmd.AddCommand(blockCloseCmd)

	carapace.Gen(blockCloseCmd).PositionalAnyCompletion(
		action.ActionBlocks(blockCloseCmd),
	)
}
