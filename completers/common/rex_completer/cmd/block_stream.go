package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var blockStreamCmd = &cobra.Command{
	Use:   "stream BLOCK [DATA_REF]",
	Short: "Stream block data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockStreamCmd).Standalone()

	blockCmd.AddCommand(blockStreamCmd)

	carapace.Gen(blockStreamCmd).PositionalCompletion(
		action.ActionBlocks(blockStreamCmd),
	)
}
