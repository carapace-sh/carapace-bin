package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var zoomCmd = &cobra.Command{
	Use:   "zoom [BLOCK]",
	Short: "Zoom a block to fill its window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(zoomCmd).Standalone()

	zoomCmd.Flags().Bool("off", false, "disable zoom")
	zoomCmd.Flags().Bool("on", false, "enable zoom")

	common.AddScopeFlags(zoomCmd)

	rootCmd.AddCommand(zoomCmd)

	carapace.Gen(zoomCmd).PositionalCompletion(
		action.ActionBlocks(zoomCmd),
	)
}
