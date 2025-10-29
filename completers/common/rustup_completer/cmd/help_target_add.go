package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_target_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a target to a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_target_addCmd).Standalone()

	help_targetCmd.AddCommand(help_target_addCmd)
}
