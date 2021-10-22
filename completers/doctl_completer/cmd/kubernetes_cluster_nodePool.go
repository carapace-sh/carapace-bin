package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_nodePoolCmd = &cobra.Command{
	Use:   "node-pool",
	Short: "Display commands for managing node pools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_nodePoolCmd).Standalone()
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_nodePoolCmd)
}
