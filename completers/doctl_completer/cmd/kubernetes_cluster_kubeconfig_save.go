package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_kubeconfig_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a cluster's credentials to your local kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_kubeconfig_saveCmd).Standalone()
	kubernetes_cluster_kubeconfig_saveCmd.Flags().Int("expiry-seconds", 0, "The length of time the cluster credentials will be valid for in seconds. By default, the credentials are automatically renewed as needed.")
	kubernetes_cluster_kubeconfig_saveCmd.Flags().Bool("set-current-context", true, "Boolean indicating whether to set the current kubectl context to that of the new cluster")
	kubernetes_cluster_kubeconfigCmd.AddCommand(kubernetes_cluster_kubeconfig_saveCmd)
}
