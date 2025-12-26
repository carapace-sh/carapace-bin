package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdfsig",
	Short: "Portable Document Format (PDF) digital signatures tool",
	Long:  "https://man.archlinux.org/man/pdfsig.1",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "print usage information")
	rootCmd.Flags().BoolS("add-signature", "add-signature", false, "adds a new signature to the document")
	rootCmd.Flags().BoolS("aia", "aia", false, "use Authority Information Access (AIA) extension for certificate fetching")
	rootCmd.Flags().StringS("assert-signer", "assert-signer", "", "This option checks whether the signature covering the full document been made with the specified key. The key is either specified as a fingerprint or a file listing fingerprints. The fingerprint must be given or listed in compact format (no colons or spaces in between). If fpr_or_file specifies a file, empty lines are ignored as well as all lines starting with a hash sign. Only available for GnuPG backend.")
	rootCmd.Flags().StringS("backend", "backend", "", "use given backend for signing/verification")
	rootCmd.Flags().StringS("digest", "digest", "", "name of the digest algorithm")
	rootCmd.Flags().BoolS("dump", "dump", false, "dump all signatures into current directory")
	rootCmd.Flags().BoolS("enable-pgp", "enable-pgp", false, "Enable pgp signatures in the GnuPG backend. Only available for GnuPG backend")
	rootCmd.Flags().BoolS("etsi", "etsi", false, "create a signature of type ETSI.CAdES.detached instead of adbe.pkcs7.detached")
	rootCmd.Flags().BoolS("h", "h", false, "print usage information")
	rootCmd.Flags().BoolP("help", "help", false, "print usage information")
	rootCmd.Flags().StringS("kpw", "kpw", "", "password for the signing key (might be missing if the key isn't password protected)")
	rootCmd.Flags().BoolS("list-backends", "list-backends", false, "print cryptographic signature backends")
	rootCmd.Flags().BoolS("list-nicks", "list-nicks", false, "list available nicknames in the NSS database")
	rootCmd.Flags().StringS("new-signature-field-name", "new-signature-field-name", "", "field name used for the newly added signature. A random ID will be used if empty")
	rootCmd.Flags().StringS("nick", "nick", "", "use the certificate with the given nickname/fingerprint for signing")
	rootCmd.Flags().BoolS("no-appearance", "no-appearance", false, "don't add appearance information when signing existing fields")
	rootCmd.Flags().BoolS("no-ocsp", "no-ocsp", false, "don't perform online OCSP certificate revocation check")
	rootCmd.Flags().BoolS("nocert", "nocert", false, "don't perform certificate validation")
	rootCmd.Flags().StringS("nss-pwd", "nss-pwd", "", "password to access the NSS database (if any)")
	rootCmd.Flags().StringS("nssdir", "nssdir", "", "path to directory of libnss3 database")
	rootCmd.Flags().StringS("opw", "opw", "", "owner password (for encrypted files)")
	rootCmd.Flags().StringS("reason", "reason", "", "reason for signing")
	rootCmd.Flags().StringS("sign", "sign", "", "sign the document in the given signature field (by name or number)")
	rootCmd.Flags().StringS("upw", "upw", "", "user password (for encrypted files)")
	rootCmd.Flags().BoolS("v", "v", false, "print copyright and version info")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"nssdir": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
