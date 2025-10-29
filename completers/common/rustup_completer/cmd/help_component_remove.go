package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_component_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a component from a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_component_removeCmd).Standalone()

	help_componentCmd.AddCommand(help_component_removeCmd)
}
