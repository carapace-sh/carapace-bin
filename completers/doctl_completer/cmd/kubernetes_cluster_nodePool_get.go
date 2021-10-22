package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_nodePool_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about a cluster's node pool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_nodePool_getCmd).Standalone()
	kubernetes_cluster_nodePool_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Size`, `Count`, `Tags`, `Labels`, `Taints`, `Nodes`")
	kubernetes_cluster_nodePool_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	kubernetes_cluster_nodePoolCmd.AddCommand(kubernetes_cluster_nodePool_getCmd)
}
