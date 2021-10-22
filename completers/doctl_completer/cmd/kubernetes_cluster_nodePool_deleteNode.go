package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_nodePool_deleteNodeCmd = &cobra.Command{
	Use:   "delete-node",
	Short: "Delete a node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_nodePool_deleteNodeCmd).Standalone()
	kubernetes_cluster_nodePool_deleteNodeCmd.Flags().BoolP("force", "f", false, "Delete the node without a confirmation prompt")
	kubernetes_cluster_nodePool_deleteNodeCmd.Flags().Bool("skip-drain", false, "Skip draining the node before deletion")
	kubernetes_cluster_nodePoolCmd.AddCommand(kubernetes_cluster_nodePool_deleteNodeCmd)
}
