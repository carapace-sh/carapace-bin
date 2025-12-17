package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubeadm"
	"github.com/spf13/cobra"
)

var certs_checkExpirationCmd = &cobra.Command{
	Use:   "check-expiration",
	Short: "Check certificates expiration for a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_checkExpirationCmd).Standalone()

	certs_checkExpirationCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	certs_checkExpirationCmd.Flags().String("cert-dir", "", "The path where to save the certificates")
	certs_checkExpirationCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	certs_checkExpirationCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	certs_checkExpirationCmd.Flags().StringP("output", "o", "", "Output format. One of: text|json|yaml|kyaml|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file.")
	certs_checkExpirationCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	certsCmd.AddCommand(certs_checkExpirationCmd)

	carapace.Gen(certs_checkExpirationCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
		"output":     kubeadm.ActionOutputFormats(),
	})
}
