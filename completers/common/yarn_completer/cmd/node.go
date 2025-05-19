package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:                "node",
	Short:              "Run node with the hook already setup",
	GroupID:            "general",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(nodeCmd).Standalone()

	rootCmd.AddCommand(nodeCmd)

	carapace.Gen(nodeCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("node"),
	)
}
