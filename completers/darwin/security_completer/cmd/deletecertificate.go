package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteCertificateCmd = &cobra.Command{
	Use:   "delete-certificate",
	Short: "Delete a certificate from a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCertificateCmd).Standalone()
	deleteCertificateCmd.Flags().StringP("hash", "Z", "", "Specify certificate to delete by SHA-1 hash")
	deleteCertificateCmd.Flags().StringP("name", "c", "", "Specify certificate to delete by common name")
	deleteCertificateCmd.Flags().BoolP("trust", "t", false, "Also delete user trust settings for this certificate")
	rootCmd.AddCommand(deleteCertificateCmd)
}
