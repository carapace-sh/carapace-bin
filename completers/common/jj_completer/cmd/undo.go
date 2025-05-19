package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:                "undo",
	Short:              "Undo an operation (shortcut for `jj op undo`)",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(undoCmd).Standalone()

	rootCmd.AddCommand(undoCmd)

	carapace.Gen(undoCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("jj", "op", "undo"),
	)
}
