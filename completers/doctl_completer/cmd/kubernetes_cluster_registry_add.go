package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_registry_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add container registry support to Kubernetes clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_registry_addCmd).Standalone()
	kubernetes_cluster_registryCmd.AddCommand(kubernetes_cluster_registry_addCmd)
}
