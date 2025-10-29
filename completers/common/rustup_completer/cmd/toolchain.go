package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var toolchainCmd = &cobra.Command{
	Use:   "toolchain",
	Short: "Install, uninstall, or list toolchains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolchainCmd).Standalone()

	toolchainCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(toolchainCmd)
}
