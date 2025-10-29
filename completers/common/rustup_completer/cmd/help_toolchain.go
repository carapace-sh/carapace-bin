package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_toolchainCmd = &cobra.Command{
	Use:   "toolchain",
	Short: "Install, uninstall, or list toolchains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_toolchainCmd).Standalone()

	helpCmd.AddCommand(help_toolchainCmd)
}
