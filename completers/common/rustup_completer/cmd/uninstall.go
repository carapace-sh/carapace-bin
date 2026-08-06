package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:                "uninstall",
	Short:              "Uninstall the given toolchains",
	Aliases:            []string{"remove", "rm", "delete", "del"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("rustup", "toolchain", "uninstall"),
	)
}
