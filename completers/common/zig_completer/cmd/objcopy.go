package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var objcopyCmd = &cobra.Command{
	Use:                "objcopy",
	Short:              "Manipulate executables and relocatables",
	GroupID:            "wrapper",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(objcopyCmd).Standalone()

	rootCmd.AddCommand(objcopyCmd)

	carapace.Gen(objcopyCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("objcopy"),
	)
}
