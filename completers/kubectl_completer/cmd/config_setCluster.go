package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_setClusterCmd = &cobra.Command{
	Use:   "set-cluster",
	Short: "Set a cluster entry in kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setClusterCmd).Standalone()
	config_setClusterCmd.Flags().Bool("embed-certs", false, "embed-certs for the cluster entry in kubeconfig")
	config_setClusterCmd.Flags().String("proxy-url", "", "proxy-url for the cluster entry in kubeconfig")
	config_setClusterCmd.Flag("embed-certs").NoOptDefVal = "true"
	configCmd.AddCommand(config_setClusterCmd)
}
