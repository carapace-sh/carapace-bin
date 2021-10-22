package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Display commands for managing your local kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_kubeconfigCmd).Standalone()
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_kubeconfigCmd)
}
