package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rustup_completer/cmd/action"
	"github.com/spf13/cobra"
)

var toolchain_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_uninstallCmd).Standalone()

	toolchain_uninstallCmd.Flags().BoolP("help", "h", false, "Prints help information")
	toolchainCmd.AddCommand(toolchain_uninstallCmd)

	carapace.Gen(toolchain_uninstallCmd).PositionalAnyCompletion(
		action.ActionToolchains().FilterArgs(),
	)
}
