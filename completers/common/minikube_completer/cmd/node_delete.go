package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var node_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a node from a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_deleteCmd).Standalone()
	nodeCmd.AddCommand(node_deleteCmd)
}
