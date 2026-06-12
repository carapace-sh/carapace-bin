package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var ranlibCmd = &cobra.Command{
	Use:                "ranlib",
	Short:              "Use Zig as a drop-in ranlib",
	GroupID:            "wrapper",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(ranlibCmd).Standalone()

	rootCmd.AddCommand(ranlibCmd)

	carapace.Gen(ranlibCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("ranlib"),
	)
}
