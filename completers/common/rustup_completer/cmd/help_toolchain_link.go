package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_toolchain_linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Create a custom toolchain by symlinking to a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_toolchain_linkCmd).Standalone()

	help_toolchainCmd.AddCommand(help_toolchain_linkCmd)
}
