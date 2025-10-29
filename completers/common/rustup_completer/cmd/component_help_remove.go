package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var component_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a component from a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(component_help_removeCmd).Standalone()

	component_helpCmd.AddCommand(component_help_removeCmd)
}
