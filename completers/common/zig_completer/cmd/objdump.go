package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var objdumpCmd = &cobra.Command{
	Use:                "objdump",
	Short:              "Print information about executables and relocatables",
	GroupID:            "wrapper",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(objdumpCmd).Standalone()

	rootCmd.AddCommand(objdumpCmd)

	carapace.Gen(objdumpCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("objdump"),
	)
}
