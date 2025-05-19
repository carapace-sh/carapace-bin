package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certs_renew_superAdminConfCmd = &cobra.Command{
	Use:   "super-admin.conf",
	Short: "Renew the certificate embedded in the kubeconfig file for the super-admin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_renew_superAdminConfCmd).Standalone()

	certs_renew_superAdminConfCmd.Flags().String("cert-dir", "", "The path where to save the certificates")
	certs_renew_superAdminConfCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_renew_superAdminConfCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	certs_renewCmd.AddCommand(certs_renew_superAdminConfCmd)

	carapace.Gen(certs_renew_superAdminConfCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
