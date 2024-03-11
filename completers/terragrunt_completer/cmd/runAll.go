package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var runAllCmd = &cobra.Command{
	Use:                "run-all",
	Short:              "Run a terraform command against a 'stack' by running the specified command in each subfolder",
	GroupID:            "terragrunt",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(runAllCmd).Standalone()

	rootCmd.AddCommand(runAllCmd)

	carapace.Gen(runAllCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("terraform"),
	)
}
