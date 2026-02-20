package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Show a tree of dependencies for a specific global environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_treeCmd).Standalone()

	global_helpCmd.AddCommand(global_help_treeCmd)
}
