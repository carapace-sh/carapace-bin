package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_nodePool_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing node pool in a cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_nodePool_updateCmd).Standalone()
	kubernetes_cluster_nodePool_updateCmd.Flags().Bool("auto-scale", false, "Boolean indicating whether to enable auto-scaling on the node pool")
	kubernetes_cluster_nodePool_updateCmd.Flags().Int("count", 0, "The size of (number of nodes in) the node pool")
	kubernetes_cluster_nodePool_updateCmd.Flags().StringSlice("label", []string{}, "Label in key=value notation to apply to the node pool, repeat to add multiple labels at once. Omitted labels will be removed from the node pool if the flag is specified.")
	kubernetes_cluster_nodePool_updateCmd.Flags().Int("max-nodes", 0, "Maximum number of nodes in the node pool when autoscaling is enabled")
	kubernetes_cluster_nodePool_updateCmd.Flags().Int("min-nodes", 0, "Minimum number of nodes in the node pool when autoscaling is enabled")
	kubernetes_cluster_nodePool_updateCmd.Flags().String("name", "", "Name of the node pool")
	kubernetes_cluster_nodePool_updateCmd.Flags().StringSlice("tag", []string{}, "Tag to apply to the node pool; you can supply multiple `--tag` arguments to specify additional tags. Omitted tags will be removed from the node pool if the flag is specified.")
	kubernetes_cluster_nodePool_updateCmd.Flags().StringSlice("taint", []string{}, "Taint in key[=value:]effect notation to apply to the node pool, repeat to add multiple taints at once. Omitted taints will be removed from the node pool if the flag is specified.")
	kubernetes_cluster_nodePoolCmd.AddCommand(kubernetes_cluster_nodePool_updateCmd)
}
