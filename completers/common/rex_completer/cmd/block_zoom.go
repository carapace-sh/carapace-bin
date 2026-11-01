package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var blockZoomCmd = &cobra.Command{
	Use:   "zoom [BLOCK]",
	Short: "Zoom a block to fill its window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockZoomCmd).Standalone()

	blockZoomCmd.Flags().Bool("off", false, "disable zoom")
	blockZoomCmd.Flags().Bool("on", false, "enable zoom")

	common.AddScopeFlags(blockZoomCmd)

	blockCmd.AddCommand(blockZoomCmd)

	carapace.Gen(blockZoomCmd).PositionalCompletion(
		action.ActionBlocks(blockZoomCmd),
	)
}
