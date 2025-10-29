package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_component_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed and available components",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_component_listCmd).Standalone()

	help_componentCmd.AddCommand(help_component_listCmd)
}
