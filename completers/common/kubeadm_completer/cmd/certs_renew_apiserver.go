package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certs_renew_apiserverCmd = &cobra.Command{
	Use:   "apiserver",
	Short: "Renew the certificate for serving the Kubernetes API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_renew_apiserverCmd).Standalone()

	certs_renew_apiserverCmd.Flags().String("cert-dir", "", "The path where to save the certificates")
	certs_renew_apiserverCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_renew_apiserverCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	certs_renewCmd.AddCommand(certs_renew_apiserverCmd)

	carapace.Gen(certs_renew_apiserverCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
