package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_nodePool_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List a cluster's node pools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_nodePool_listCmd).Standalone()
	kubernetes_cluster_nodePool_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Size`, `Count`, `Tags`, `Labels`, `Taints`, `Nodes`")
	kubernetes_cluster_nodePool_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	kubernetes_cluster_nodePoolCmd.AddCommand(kubernetes_cluster_nodePool_listCmd)
}
