package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_nodePool_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a node pool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_nodePool_deleteCmd).Standalone()
	kubernetes_cluster_nodePool_deleteCmd.Flags().BoolP("force", "f", false, "Delete node pool without confirmation prompt")
	kubernetes_cluster_nodePoolCmd.AddCommand(kubernetes_cluster_nodePool_deleteCmd)
}
