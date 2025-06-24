package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemd-analyze_completer/cmd/action"
	"github.com/spf13/cobra"
)

var criticalChainCmd = &cobra.Command{
	Use:   "critical-chain",
	Short: "Print a tree of the time critical chain of units",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(criticalChainCmd).Standalone()

	rootCmd.AddCommand(criticalChainCmd)

	carapace.Gen(criticalChainCmd).PositionalAnyCompletion(
		action.ActionUnits(criticalChainCmd).FilterArgs(),
	)
}
