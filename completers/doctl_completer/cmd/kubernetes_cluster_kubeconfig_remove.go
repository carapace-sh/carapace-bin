package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_kubeconfig_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a cluster's credentials from your local kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_kubeconfig_removeCmd).Standalone()
	kubernetes_cluster_kubeconfigCmd.AddCommand(kubernetes_cluster_kubeconfig_removeCmd)
}
