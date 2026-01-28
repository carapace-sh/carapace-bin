package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var reqCmd = &cobra.Command{
	Use:     "req",
	Short:   "PKCS#10 X.509 Certificate Signing Request (CSR) Management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reqCmd).Standalone()

	reqCmd.Flags().StringS("CA", "CA", "", "Issuer cert to use for signing a cert, implies -x509")
	reqCmd.Flags().StringS("CAkey", "CAkey", "", "Issuer private key to use with -CA; default is -CA arg (Required by some CA's)")
	reqCmd.Flags().StringS("addext", "addext", "", "Additional cert extension key=value pair (may be given more than once)")
	reqCmd.Flags().BoolS("batch", "batch", false, "Do not ask anything during request generation")
	reqCmd.Flags().StringS("cipher", "cipher", "", "Specify the cipher for private key encryption")
	reqCmd.Flags().StringS("config", "config", "", "Request template file")
	reqCmd.Flags().StringS("copy_extensions", "copy_extensions", "", "copy extensions from request when using -x509")
	reqCmd.Flags().StringS("days", "days", "", "Number of days certificate is valid for")
	reqCmd.Flags().StringS("extensions", "extensions", "", "Cert or request extension section (override value in config file)")
	reqCmd.Flags().StringS("in", "in", "", "X.509 request input file (default stdin)")
	reqCmd.Flags().StringS("inform", "inform", "", "CSR input format to use (PEM or DER; by default try PEM first)")
	reqCmd.Flags().StringS("key", "key", "", "Key for signing, and to include unless -in given")
	reqCmd.Flags().StringS("keyform", "keyform", "", "Key file format (DER/PEM)")
	reqCmd.Flags().StringS("keyout", "keyout", "", "File to write private key to")
	reqCmd.Flags().BoolS("modulus", "modulus", false, "RSA modulus")
	reqCmd.Flags().BoolS("multivalue-rdn", "multivalue-rdn", false, "Deprecated; multi-valued RDNs support is always on.")
	reqCmd.Flags().StringSliceS("nameopt", "nameopt", nil, "Certificate subject/issuer name printing options")
	reqCmd.Flags().BoolS("new", "new", false, "New request")
	reqCmd.Flags().BoolS("newhdr", "newhdr", false, "Output \"NEW\" in the header lines")
	reqCmd.Flags().StringS("newkey", "newkey", "", "Generate new key with [<alg>:]<nbits> or <alg>[:<file>] or param:<file>")
	reqCmd.Flags().BoolS("nodes", "nodes", false, "Don't encrypt private keys; deprecated")
	reqCmd.Flags().BoolS("noenc", "noenc", false, "Don't encrypt private keys")
	reqCmd.Flags().BoolS("noout", "noout", false, "Do not output REQ")
	reqCmd.Flags().StringS("not_after", "not_after", "", "[CC]YYMMDDHHMMSSZ value for notAfter certificate field, overrides -days")
	reqCmd.Flags().StringS("not_before", "not_before", "", "[CC]YYMMDDHHMMSSZ value for notBefore certificate field")
	reqCmd.Flags().StringS("out", "out", "", "Output file")
	reqCmd.Flags().StringS("outform", "outform", "", "Output format - DER or PEM")
	reqCmd.Flags().StringS("passin", "passin", "", "Private key and certificate password source")
	reqCmd.Flags().StringS("passout", "passout", "", "Output file pass phrase source")
	reqCmd.Flags().StringSliceS("pkeyopt", "pkeyopt", nil, "Public key options as opt:value")
	reqCmd.Flags().BoolS("precert", "precert", false, "Add a poison extension to generated cert (implies -new)")
	reqCmd.Flags().BoolS("pubkey", "pubkey", false, "Output public key")
	reqCmd.Flags().BoolS("quiet", "quiet", false, "Terse output")
	reqCmd.Flags().StringS("reqexts", "reqexts", "", "An alias for -extensions")
	reqCmd.Flags().StringSliceS("reqopt", "reqopt", nil, "Various request text options")
	reqCmd.Flags().StringS("section", "section", "", "Config section to use (default \"req\")")
	reqCmd.Flags().StringS("set_serial", "set_serial", "", "Serial number to use")
	reqCmd.Flags().StringSliceS("sigopt", "sigopt", nil, "Signature parameter in n:v form")
	reqCmd.Flags().StringS("subj", "subj", "", "Set or modify subject of request or cert")
	reqCmd.Flags().BoolS("subject", "subject", false, "Print the subject of the output request or cert")
	reqCmd.Flags().BoolS("text", "text", false, "Text form of request")
	reqCmd.Flags().BoolS("utf8", "utf8", false, "Input characters are UTF8 (default ASCII)")
	reqCmd.Flags().BoolS("verbose", "verbose", false, "Verbose output")
	reqCmd.Flags().BoolS("verify", "verify", false, "Verify self-signature on the request")
	reqCmd.Flags().StringSliceS("vfyopt", "vfyopt", nil, "Verification parameter in n:v form")
	reqCmd.Flags().BoolS("x509", "x509", false, "Output an X.509 certificate structure instead of a cert request")
	reqCmd.Flags().BoolS("x509v1", "x509v1", false, "Request cert generation with X.509 version 1")
	common.AddProviderFlags(reqCmd)
	common.AddRandomStateFlags(reqCmd)
	rootCmd.AddCommand(reqCmd)

	reqCmd.MarkFlagsMutuallyExclusive("reqexts", "extensions")

	carapace.Gen(reqCmd).FlagCompletion(carapace.ActionMap{
		"CA":      carapace.ActionFiles(),
		"CAkey":   carapace.ActionFiles(),
		"cipher":  action.ActionCipherAlgorithms(reqCmd),
		"config":  carapace.ActionFiles(),
		"in":      carapace.ActionFiles(),
		"inform":  carapace.ActionValues("DER", "PEM"),
		"key":     carapace.ActionFiles(),
		"keyform": carapace.ActionValues("DER", "PEM", "P12"),
		"out":     carapace.ActionFiles(),
		"outform": carapace.ActionValues("DER", "PEM"),
	})
}
