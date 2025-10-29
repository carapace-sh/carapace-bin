package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var show_activeToolchainCmd = &cobra.Command{
	Use:   "active-toolchain",
	Short: "Show the active toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_activeToolchainCmd).Standalone()

	show_activeToolchainCmd.Flags().BoolP("help", "h", false, "Print help")
	show_activeToolchainCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output with rustc information")
	showCmd.AddCommand(show_activeToolchainCmd)
}
