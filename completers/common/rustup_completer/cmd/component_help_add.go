package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var component_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a component to a Rust toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(component_help_addCmd).Standalone()

	component_helpCmd.AddCommand(component_help_addCmd)
}
