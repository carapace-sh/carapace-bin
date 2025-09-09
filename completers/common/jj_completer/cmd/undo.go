package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:                "undo",
	Short:              "Undo the last operation",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(undoCmd).Standalone()

	undoCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(undoCmd)

	carapace.Gen(undoCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("jj", "op", "undo"),
	)
}
