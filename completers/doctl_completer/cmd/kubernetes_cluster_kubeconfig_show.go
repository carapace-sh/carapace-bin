package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_kubeconfig_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a Kubernetes cluster's kubeconfig YAML",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_kubeconfig_showCmd).Standalone()
	kubernetes_cluster_kubeconfig_showCmd.Flags().Int("expiry-seconds", 0, "The length of time the cluster credentials will be valid for in seconds. By default, the credentials expire after seven days.")
	kubernetes_cluster_kubeconfigCmd.AddCommand(kubernetes_cluster_kubeconfig_showCmd)
}
