package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mkcert",
	Short: "simple tool for making locally-trusted development certificates",
	Long:  "https://github.com/FiloSottile/mkcert",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("CAROOT", "CAROOT", false, "Print the CA certificate and key storage location")
	rootCmd.Flags().StringS("cert-file", "cert-file", "", "Customize the output paths")
	rootCmd.Flags().BoolS("client", "client", false, "Generate a certificate for client authentication")
	rootCmd.Flags().StringS("csr", "csr", "", "Generate a certificate based on the supplied CSR")
	rootCmd.Flags().BoolS("ecdsa", "ecdsa", false, "Generate a certificate with an ECDSA key")
	rootCmd.Flags().StringS("install", "install", "", "Install the local CA in the system trust store")
	rootCmd.Flags().StringS("key-file", "key-file", "", "Customize the output paths")
	rootCmd.Flags().StringS("p12-file", "p12-file", "", "Customize the output paths")
	rootCmd.Flags().BoolS("pkcs12", "pkcs12", false, "Generate a \".p12\" PKCS #12 file, also know as a \".pfx\" file")
	rootCmd.Flags().StringS("uninstall", "uninstall", "", "Uninstall the local CA (but do not delete it)")

	for _, f := range []string{
		"CAROOT",
		"client",
		"csr",
		"ecdsa",
		"key-file",
		"p12-file",
		"pkcs12",
		"uninstall",
	} {
		rootCmd.MarkFlagsMutuallyExclusive("csr", f)
	}

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cert-file": carapace.ActionFiles(),
		"csr":       carapace.ActionFiles(),
		"key-file":  carapace.ActionFiles(),
		"p12-file":  carapace.ActionFiles(),
	})
}
