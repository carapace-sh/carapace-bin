package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var ccCmd = &cobra.Command{
	Use:                "cc",
	Short:              "Use Zig as a drop-in C compiler",
	Run:                func(cmd *cobra.Command, args []string) {},
	GroupID:            "wrapper",
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(ccCmd).Standalone()

	rootCmd.AddCommand(ccCmd)

	carapace.Gen(ccCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("cc"),
	)
}
