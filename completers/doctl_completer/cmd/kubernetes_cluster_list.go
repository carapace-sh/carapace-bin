package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve the list of Kubernetes clusters for your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_listCmd).Standalone()
	kubernetes_cluster_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Region`, `Version`, `AutoUpgrade`, `HAControlPlane`, `Status`, `Endpoint`, `IPv4`, `ClusterSubnet`, `ServiceSubnet`, `Tags`, `Created`, `Updated`, `NodePools`")
	kubernetes_cluster_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_listCmd)
}
