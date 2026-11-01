package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:                "install",
	Short:              "Install or update the given toolchains, or by default the active toolchain",
	Aliases:            []string{"add"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(installCmd).Standalone()

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("rustup", "toolchain", "install"),
	)
}
