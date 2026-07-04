package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import items into a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()
	importCmd.Flags().BoolP("allow-all", "A", false, "Allow any application to access imported key without warning")
	importCmd.Flags().StringP("attribute", "a", "", "Specify optional extended attribute name and value")
	importCmd.Flags().StringP("format", "f", "", "Format: openssl, bsafe, raw, pkcs7, pkcs8, pkcs12, x509, etc.")
	importCmd.Flags().StringP("keychain", "k", "", "Specify keychain into which item(s) will be imported")
	importCmd.Flags().BoolP("non-extractable", "x", false, "Private keys are non-extractable after being imported")
	importCmd.Flags().StringP("passphrase", "P", "", "Specify unwrapping passphrase immediately")
	importCmd.Flags().StringP("tool", "T", "", "Specify an application which may access imported key")
	importCmd.Flags().StringP("type", "t", "", "Type: cert, pub, priv, session, agg")
	importCmd.Flags().BoolP("wrap", "w", false, "Private keys are wrapped and must be unwrapped on import")
	rootCmd.AddCommand(importCmd)
	carapace.Gen(importCmd).PositionalCompletion(carapace.ActionFiles())
	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"format":   carapace.ActionValues("openssl", "bsafe", "raw", "pkcs7", "pkcs8", "pkcs12", "x509", "openssh1", "openssh2", "pemseq"),
		"keychain": security.ActionKeychains(),
		"type":     carapace.ActionValues("cert", "pub", "priv", "session", "agg"),
	})
}
