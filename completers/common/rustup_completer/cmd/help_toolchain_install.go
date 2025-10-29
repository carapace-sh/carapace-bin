package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_toolchain_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install or update the given toolchains, or by default the active toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_toolchain_installCmd).Standalone()

	help_toolchainCmd.AddCommand(help_toolchain_installCmd)
}
