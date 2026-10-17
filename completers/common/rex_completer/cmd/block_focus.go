package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var blockFocusCmd = &cobra.Command{
	Use:   "focus [BLOCK]",
	Short: "Focus a block",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockFocusCmd).Standalone()

	common.AddWindowScopeFlags(blockFocusCmd)

	blockCmd.AddCommand(blockFocusCmd)

	carapace.Gen(blockFocusCmd).PositionalCompletion(
		action.ActionBlocks(blockFocusCmd),
	)
}
