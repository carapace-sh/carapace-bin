package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:                "fix",
	Short:              "apply fixes suggested by static checkers",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(fixCmd).Standalone()

	rootCmd.AddCommand(fixCmd)

	carapace.Gen(fixCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("go", "tool", "fix"),
	)
}
