package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var blockAttachCmd = &cobra.Command{
	Use:   "attach BLOCK",
	Short: "Attach to a block",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockAttachCmd).Standalone()

	blockCmd.AddCommand(blockAttachCmd)

	carapace.Gen(blockAttachCmd).PositionalCompletion(
		action.ActionBlocks(blockAttachCmd),
	)
}
