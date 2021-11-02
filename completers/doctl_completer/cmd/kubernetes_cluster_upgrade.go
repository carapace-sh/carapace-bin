package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades a cluster to a new Kubernetes version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_upgradeCmd).Standalone()
	kubernetes_cluster_upgradeCmd.Flags().String("version", "latest", "The desired Kubernetes version. Possible values: see `doctl k8s get-upgrades <cluster>`.")
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_upgradeCmd)
}
