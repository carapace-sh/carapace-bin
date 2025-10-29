package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_component_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a component to a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_component_addCmd).Standalone()

	help_componentCmd.AddCommand(help_component_addCmd)
}
