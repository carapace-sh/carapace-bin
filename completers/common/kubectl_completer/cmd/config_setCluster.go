package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_setClusterCmd = &cobra.Command{
	Use:   "set-cluster NAME [--server=server] [--certificate-authority=path/to/certificate/authority] [--insecure-skip-tls-verify=true] [--tls-server-name=example.com]",
	Short: "Set a cluster entry in kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setClusterCmd).Standalone()

	config_setClusterCmd.Flags().String("certificate-authority", "", "Path to certificate-authority file for the cluster entry in kubeconfig")
	config_setClusterCmd.Flags().String("embed-certs", "", "embed-certs for the cluster entry in kubeconfig")
	config_setClusterCmd.Flags().String("insecure-skip-tls-verify", "", "insecure-skip-tls-verify for the cluster entry in kubeconfig")
	config_setClusterCmd.Flags().String("proxy-url", "", "proxy-url for the cluster entry in kubeconfig")
	config_setClusterCmd.Flags().String("server", "", "server for the cluster entry in kubeconfig")
	config_setClusterCmd.Flags().String("tls-server-name", "", "tls-server-name for the cluster entry in kubeconfig")
	config_setClusterCmd.Flag("embed-certs").NoOptDefVal = " "
	config_setClusterCmd.Flag("insecure-skip-tls-verify").NoOptDefVal = " "
	configCmd.AddCommand(config_setClusterCmd)

	carapace.Gen(config_setClusterCmd).FlagCompletion(carapace.ActionMap{
		"certificate-authority": carapace.ActionFiles(),
	})
}
