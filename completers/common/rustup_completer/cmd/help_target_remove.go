package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_target_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a target from a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_target_removeCmd).Standalone()

	help_targetCmd.AddCommand(help_target_removeCmd)
}
