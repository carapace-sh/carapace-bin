package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rcCmd = &cobra.Command{
	Use:                "rc",
	Short:              "Use Zig as a drop-in rc.exe",
	GroupID:            "wrapper",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(rcCmd).Standalone()

	rootCmd.AddCommand(rcCmd)

	carapace.Gen(rcCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("rc"),
	)
}
