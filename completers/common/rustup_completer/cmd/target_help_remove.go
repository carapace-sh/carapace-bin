package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var target_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a target from a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(target_help_removeCmd).Standalone()

	target_helpCmd.AddCommand(target_help_removeCmd)
}
