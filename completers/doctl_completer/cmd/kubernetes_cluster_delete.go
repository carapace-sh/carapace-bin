package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Kubernetes clusters ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_deleteCmd).Standalone()
	kubernetes_cluster_deleteCmd.Flags().Bool("dangerous", false, "Boolean indicating whether to delete the cluster's associated resources like load balancers, volumes and volume snapshots")
	kubernetes_cluster_deleteCmd.Flags().BoolP("force", "f", false, "Boolean indicating whether to delete the cluster without a confirmation prompt")
	kubernetes_cluster_deleteCmd.Flags().Bool("update-kubeconfig", true, "Boolean indicating whether to remove the deleted cluster from your kubeconfig")
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_deleteCmd)
}
