package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var blockRenameCmd = &cobra.Command{
	Use:   "rename [BLOCK] LABEL",
	Short: "Rename a block",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockRenameCmd).Standalone()

	common.AddWindowScopeFlags(blockRenameCmd)

	blockCmd.AddCommand(blockRenameCmd)

	carapace.Gen(blockRenameCmd).PositionalCompletion(
		action.ActionBlocks(blockRenameCmd),
	)
}
