package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_nodePool_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new node pool for a cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_nodePool_createCmd).Standalone()
	kubernetes_cluster_nodePool_createCmd.Flags().Bool("auto-scale", false, "Boolean indicating whether to enable auto-scaling on the node pool")
	kubernetes_cluster_nodePool_createCmd.Flags().Int("count", 0, "The size of (number of nodes in) the node pool (required)")
	kubernetes_cluster_nodePool_createCmd.Flags().StringSlice("label", []string{}, "Label in key=value notation to apply to the node pool; repeat to specify additional labels. An existing label is removed from the node pool if it is not specified by any flag.")
	kubernetes_cluster_nodePool_createCmd.Flags().Int("max-nodes", 0, "Maximum number of nodes in the node pool when autoscaling is enabled")
	kubernetes_cluster_nodePool_createCmd.Flags().Int("min-nodes", 0, "Minimum number of nodes in the node pool when autoscaling is enabled")
	kubernetes_cluster_nodePool_createCmd.Flags().String("name", "", "Name of the node pool (required)")
	kubernetes_cluster_nodePool_createCmd.Flags().String("size", "", "Size of the nodes in the node pool (To see possible values: call `doctl kubernetes options sizes`) (required)")
	kubernetes_cluster_nodePool_createCmd.Flags().StringSlice("tag", []string{}, "Tag to apply to the node pool; repeat to specify additional tags. An existing tag is removed from the node pool if it is not specified by any flag.")
	kubernetes_cluster_nodePool_createCmd.Flags().StringSlice("taint", []string{}, "Taint in key[=value:]effect notation to apply to the node pool; repeat to specify additional taints. Set to the empty string \"\" to clear all taints. An existing taint is removed from the node pool if it is not specified by any flag.")
	kubernetes_cluster_nodePoolCmd.AddCommand(kubernetes_cluster_nodePool_createCmd)
}
