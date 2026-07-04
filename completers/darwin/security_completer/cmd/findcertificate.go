package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findCertificateCmd = &cobra.Command{
	Use:   "find-certificate",
	Short: "Find a certificate item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findCertificateCmd).Standalone()
	findCertificateCmd.Flags().BoolP("all", "a", false, "Find all matching certificates, not just the first")
	findCertificateCmd.Flags().StringP("email", "e", "", "Match on email address when searching")
	findCertificateCmd.Flags().BoolP("email-addresses", "m", false, "Show the email addresses in the certificate")
	findCertificateCmd.Flags().BoolP("hash", "Z", false, "Print SHA-1 hash of the certificate")
	findCertificateCmd.Flags().StringP("name", "c", "", "Match on name when searching")
	findCertificateCmd.Flags().BoolP("pem", "p", false, "Output certificate in PEM format")
	rootCmd.AddCommand(findCertificateCmd)
}
