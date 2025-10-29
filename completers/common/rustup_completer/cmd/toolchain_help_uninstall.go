package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toolchain_help_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the given toolchains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchain_help_uninstallCmd).Standalone()

	toolchain_helpCmd.AddCommand(toolchain_help_uninstallCmd)
}
