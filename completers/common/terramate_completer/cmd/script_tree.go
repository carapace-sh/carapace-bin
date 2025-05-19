package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var script_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Show a tree of all scripts in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(script_treeCmd).Standalone()

	scriptCmd.AddCommand(script_treeCmd)
}
