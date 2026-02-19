package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var global_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Show a tree of dependencies for a specific global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_treeCmd).Standalone()

	global_treeCmd.Flags().StringP("environment", "e", "", "The environment to show the tree for")
	global_treeCmd.Flags().BoolP("invert", "i", false, "Invert the dependency tree")
	globalCmd.AddCommand(global_treeCmd)

	carapace.Gen(global_treeCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionGlobalEnvironments(),
	})
}
