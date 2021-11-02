package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_getUpgradesCmd = &cobra.Command{
	Use:   "get-upgrades",
	Short: "Retrieve a list of available Kubernetes version upgrades",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_getUpgradesCmd).Standalone()
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_getUpgradesCmd)
}
