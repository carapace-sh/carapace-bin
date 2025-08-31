package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var crlCmd = &cobra.Command{
	Use:     "crl",
	Short:   "Certificate Revocation List (CRL) Management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crlCmd).Standalone()

	crlCmd.Flags().StringS("CAfile", "CAfile", "", "Verify CRL using certificates in file name")
	crlCmd.Flags().StringS("CApath", "CApath", "", "Verify CRL using certificates in dir")
	crlCmd.Flags().StringS("CAstore", "CAstore", "", "Verify CRL using certificates in store URI")
	crlCmd.Flags().BoolS("badsig", "badsig", false, "Corrupt last byte of loaded CRL signature (for test)")
	crlCmd.Flags().BoolS("crlnumber", "crlnumber", false, "Print CRL number")
	crlCmd.Flags().StringSliceS("dateopt", "dateopt", nil, "Datetime format used for printing. (rfc_822/iso_8601). Default is rfc_822.")
	crlCmd.Flags().BoolS("fingerprint", "fingerprint", false, "Print the crl fingerprint")
	crlCmd.Flags().StringS("gendelta", "gendelta", "", "Other CRL to compare/diff to the Input one")
	crlCmd.Flags().BoolS("hash", "hash", false, "Print hash value")
	crlCmd.Flags().BoolS("hash_old", "hash_old", false, "Print old-style (MD5) hash value")
	crlCmd.Flags().StringS("in", "in", "", "Input file - default stdin")
	crlCmd.Flags().StringS("inform", "inform", "", "CRL input format (DER or PEM); has no effect")
	crlCmd.Flags().BoolS("issuer", "issuer", false, "Print issuer DN")
	crlCmd.Flags().StringS("key", "key", "", "CRL signing Private key to use")
	crlCmd.Flags().StringS("keyform", "keyform", "", "Private key file format (DER/PEM/P12); has no effect")
	crlCmd.Flags().BoolS("lastupdate", "lastupdate", false, "Set lastUpdate field")
	crlCmd.Flags().StringSliceS("nameopt", "nameopt", nil, "Certificate subject/issuer name printing options")
	crlCmd.Flags().BoolS("nextupdate", "nextupdate", false, "Set nextUpdate field")
	crlCmd.Flags().BoolS("no-CAfile", "no-CAfile", false, "Do not load the default certificates file")
	crlCmd.Flags().BoolS("no-CApath", "no-CApath", false, "Do not load certificates from the default certificates directory")
	crlCmd.Flags().BoolS("no-CAstore", "no-CAstore", false, "Do not load certificates from the default certificates store")
	crlCmd.Flags().BoolS("noout", "noout", false, "No CRL output")
	crlCmd.Flags().StringS("out", "out", "", "output file - default stdout")
	crlCmd.Flags().StringS("outform", "outform", "", "Output format - default PEM")
	crlCmd.Flags().BoolS("text", "text", false, "Print out a text format version")
	crlCmd.Flags().BoolS("verify", "verify", false, "Verify CRL signature")
	common.AddProviderFlags(crlCmd)
	rootCmd.AddCommand(crlCmd)

	carapace.Gen(crlCmd).FlagCompletion(carapace.ActionMap{
		"CAfile":   carapace.ActionFiles(),
		"CApath":   carapace.ActionDirectories(),
		"gendelta": carapace.ActionFiles(),
		"in":       carapace.ActionFiles(),
		"inform":   carapace.ActionValues("DER", "PEM"),
		"key":      carapace.ActionFiles(),
		"keyform":  carapace.ActionValues("DER", "PEM", "P12"),
		"out":      carapace.ActionFiles(),
		"outform":  carapace.ActionValues("DER", "PEM"),
	})
}
