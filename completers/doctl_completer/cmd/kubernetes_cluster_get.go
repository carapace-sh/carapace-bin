package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve details about a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_getCmd).Standalone()
	kubernetes_cluster_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Region`, `Version`, `AutoUpgrade`, `HAControlPlane`, `Status`, `Endpoint`, `IPv4`, `ClusterSubnet`, `ServiceSubnet`, `Tags`, `Created`, `Updated`, `NodePools`")
	kubernetes_cluster_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_getCmd)
}
