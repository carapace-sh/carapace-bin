package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_certificate_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new certificate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_certificate_createCmd).Standalone()
	compute_certificate_createCmd.Flags().String("certificate-chain-path", "", "The path to a full PEM-formatted trust chain between the certificate authority's certificate and your domain's SSL certificate.")
	compute_certificate_createCmd.Flags().StringSlice("dns-names", []string{}, "Comma-separated list of domains for which the certificate will be issued. The domains must be managed using DigitalOcean's DNS.")
	compute_certificate_createCmd.Flags().String("leaf-certificate-path", "", "The path to a PEM-formatted public SSL certificate.")
	compute_certificate_createCmd.Flags().String("name", "", "Certificate name (required)")
	compute_certificate_createCmd.Flags().String("private-key-path", "", "The path to a PEM-formatted private-key corresponding to the SSL certificate.")
	compute_certificate_createCmd.Flags().String("type", "", "Certificate type [custom|lets_encrypt]")
	compute_certificateCmd.AddCommand(compute_certificate_createCmd)

	carapace.Gen(compute_certificate_createCmd).FlagCompletion(carapace.ActionMap{
		"certificate-chain-path": carapace.ActionFiles(),
		"leaf-certificate-path":  carapace.ActionFiles(),
		"private-key-path":       carapace.ActionFiles(),
		"type":                   carapace.ActionValues("custom", "lets_encrypt"),
	})
}
