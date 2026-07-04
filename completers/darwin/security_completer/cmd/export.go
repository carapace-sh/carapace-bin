package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export items from a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()
	exportCmd.Flags().StringP("format", "f", "", "Format: openssl, bsafe, pkcs7, pkcs8, pkcs12, x509, etc.")
	exportCmd.Flags().StringP("keychain", "k", "", "Specify keychain from which item(s) will be exported")
	exportCmd.Flags().StringP("output", "o", "", "Write output data to outfile (default: stdout)")
	exportCmd.Flags().StringP("passphrase", "P", "", "Specify wrapping passphrase immediately")
	exportCmd.Flags().BoolP("pem-armour", "p", false, "PEM armour applied to output data")
	exportCmd.Flags().StringP("type", "t", "", "Type: certs, allKeys, pubKeys, privKeys, identities, all")
	exportCmd.Flags().BoolP("wrap", "w", false, "Private keys are to be wrapped on export")
	rootCmd.AddCommand(exportCmd)
	carapace.Gen(exportCmd).FlagCompletion(carapace.ActionMap{
		"format":   carapace.ActionValues("openssl", "bsafe", "pkcs7", "pkcs8", "pkcs12", "x509", "openssh1", "openssh2", "pemseq"),
		"keychain": security.ActionKeychains(),
		"output":   carapace.ActionFiles(),
		"type":     carapace.ActionValues("certs", "allKeys", "pubKeys", "privKeys", "identities", "all"),
	})
}
