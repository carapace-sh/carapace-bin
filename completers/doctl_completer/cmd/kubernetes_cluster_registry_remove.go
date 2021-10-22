package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_registry_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove container registry support from Kubernetes clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_registry_removeCmd).Standalone()
	kubernetes_cluster_registryCmd.AddCommand(kubernetes_cluster_registry_removeCmd)
}
