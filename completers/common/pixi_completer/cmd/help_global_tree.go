package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Show a tree of dependencies for a specific global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_treeCmd).Standalone()

	help_globalCmd.AddCommand(help_global_treeCmd)
}
