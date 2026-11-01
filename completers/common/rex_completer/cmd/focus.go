package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var focusCmd = &cobra.Command{
	Use:   "focus TARGET|DIRECTION",
	Short: "Focus a block or window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(focusCmd).Standalone()

	focusCmd.Flags().Bool("wrap", false, "continue from the other side at the edge")

	common.AddWindowScopeFlags(focusCmd)

	rootCmd.AddCommand(focusCmd)

	carapace.Gen(focusCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValuesDescribed(
				"down", "focus the block below",
				"left", "focus the block to the left",
				"next", "focus the next block in layout order",
				"prev", "focus the previous block in layout order",
				"right", "focus the block to the right",
				"up", "focus the block above",
			),
			action.ActionBlocks(focusCmd),
			action.ActionWindows(focusCmd),
		).ToA(),
	)
}
