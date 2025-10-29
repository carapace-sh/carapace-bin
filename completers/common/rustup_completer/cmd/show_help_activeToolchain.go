package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var show_help_activeToolchainCmd = &cobra.Command{
	Use:   "active-toolchain",
	Short: "Show the active toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_help_activeToolchainCmd).Standalone()

	show_helpCmd.AddCommand(show_help_activeToolchainCmd)
}
