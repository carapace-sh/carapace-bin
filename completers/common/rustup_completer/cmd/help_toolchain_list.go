package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_toolchain_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed toolchains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_toolchain_listCmd).Standalone()

	help_toolchainCmd.AddCommand(help_toolchain_listCmd)
}
