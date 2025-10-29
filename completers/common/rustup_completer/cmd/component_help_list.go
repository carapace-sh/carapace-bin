package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var component_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed and available components",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(component_help_listCmd).Standalone()

	component_helpCmd.AddCommand(component_help_listCmd)
}
