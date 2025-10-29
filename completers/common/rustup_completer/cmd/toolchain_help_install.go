package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toolchain_help_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install or update the given toolchains, or by default the active toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_help_installCmd).Standalone()

	toolchain_helpCmd.AddCommand(toolchain_help_installCmd)
}
