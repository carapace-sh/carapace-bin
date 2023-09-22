package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var node_deleteCmd = &cobra.Command{
	Use:   "delete (NAME | --all)",
	Short: "Delete node(s).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_deleteCmd).Standalone()

	node_deleteCmd.Flags().BoolP("all", "a", false, "Delete all existing nodes")
	node_deleteCmd.Flags().BoolP("registries", "r", false, "Also delete registries")
	nodeCmd.AddCommand(node_deleteCmd)
}
