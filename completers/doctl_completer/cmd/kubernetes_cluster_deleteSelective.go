package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_deleteSelectiveCmd = &cobra.Command{
	Use:   "delete-selective",
	Short: "Delete a Kubernetes cluster and selectively delete resources associated with it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_deleteSelectiveCmd).Standalone()
	kubernetes_cluster_deleteSelectiveCmd.Flags().BoolP("force", "f", false, "Boolean indicating whether to delete the cluster without a confirmation prompt")
	kubernetes_cluster_deleteSelectiveCmd.Flags().StringSlice("load-balancers", []string{}, "Comma-separated list of load balancer IDs or names for deletion")
	kubernetes_cluster_deleteSelectiveCmd.Flags().StringSlice("snapshots", []string{}, "Comma-separated list of volume snapshot IDs or names for deletion")
	kubernetes_cluster_deleteSelectiveCmd.Flags().Bool("update-kubeconfig", true, "Boolean indicating whether to remove the deleted cluster from your kubeconfig")
	kubernetes_cluster_deleteSelectiveCmd.Flags().StringSlice("volumes", []string{}, "Comma-separated list of volume IDs or names for deletion")
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_deleteSelectiveCmd)
}
