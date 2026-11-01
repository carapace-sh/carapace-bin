package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certs_issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Issue a new certificate for a domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_issueCmd).Standalone()

	certs_issueCmd.Flags().String("ca", "", "CA certificate chain file")
	certs_issueCmd.Flags().Bool("challenge-only", false, "Only show challenges needed to issue a cert")
	certs_issueCmd.Flags().String("crt", "", "Certificate file")
	certs_issueCmd.Flags().String("key", "", "Certificate key file")
	certs_issueCmd.Flags().Bool("overwrite", false, "Overwrite existing certificate")

	certsCmd.AddCommand(certs_issueCmd)

	carapace.Gen(certs_issueCmd).FlagCompletion(carapace.ActionMap{
		"ca":  carapace.ActionFiles(),
		"crt": carapace.ActionFiles(),
		"key": carapace.ActionFiles(),
	})

	carapace.Gen(certs_issueCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
