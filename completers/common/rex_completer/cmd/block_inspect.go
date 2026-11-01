package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var blockInspectCmd = &cobra.Command{
	Use:   "inspect [BLOCK]",
	Short: "Inspect a block",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockInspectCmd).Standalone()

	blockInspectCmd.Flags().Bool("json", false, "output block information as JSON")

	common.AddWindowScopeFlags(blockInspectCmd)

	blockCmd.AddCommand(blockInspectCmd)

	carapace.Gen(blockInspectCmd).PositionalCompletion(
		action.ActionBlocks(blockInspectCmd),
	)
}
