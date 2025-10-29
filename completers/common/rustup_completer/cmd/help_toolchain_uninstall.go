package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_toolchain_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the given toolchains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_toolchain_uninstallCmd).Standalone()

	help_toolchainCmd.AddCommand(help_toolchain_uninstallCmd)
}
