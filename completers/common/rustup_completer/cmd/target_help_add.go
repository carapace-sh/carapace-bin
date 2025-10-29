package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var target_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a target to a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(target_help_addCmd).Standalone()

	target_helpCmd.AddCommand(target_help_addCmd)
}
