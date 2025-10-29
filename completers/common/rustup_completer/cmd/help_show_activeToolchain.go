package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_show_activeToolchainCmd = &cobra.Command{
	Use:   "active-toolchain",
	Short: "Show the active toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_show_activeToolchainCmd).Standalone()

	help_showCmd.AddCommand(help_show_activeToolchainCmd)
}
