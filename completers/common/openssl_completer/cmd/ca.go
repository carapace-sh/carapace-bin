package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var caCmd = &cobra.Command{
	Use:     "ca",
	Short:   "Certificate Authority (CA) Management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(caCmd).Standalone()

	caCmd.Flags().BoolS("batch", "batch", false, "Don't ask questions")
	caCmd.Flags().StringS("cert", "cert", "", "The CA cert")
	caCmd.Flags().StringS("certform", "certform", "", "Certificate input format (DER/PEM/P12); has no effect")
	caCmd.Flags().StringS("config", "config", "", "A config file")
	caCmd.Flags().BoolS("create_serial", "create_serial", false, "If reading serial fails, create a new random serial")
	caCmd.Flags().StringS("crl_CA_compromise", "crl_CA_compromise", "", "sets compromise time to val and the revocation reason to CACompromise")
	caCmd.Flags().StringS("crl_compromise", "crl_compromise", "", "sets compromise time to val and the revocation reason to keyCompromise")
	caCmd.Flags().StringS("crl_hold", "crl_hold", "", "the hold instruction, an OID. Sets revocation reason to certificateHold")
	caCmd.Flags().StringS("crl_lastupdate", "crl_lastupdate", "", "Sets the CRL lastUpdate time to val (YYMMDDHHMMSSZ or YYYYMMDDHHMMSSZ)")
	caCmd.Flags().StringS("crl_nextupdate", "crl_nextupdate", "", "Sets the CRL nextUpdate time to val (YYMMDDHHMMSSZ or YYYYMMDDHHMMSSZ)")
	caCmd.Flags().StringS("crl_reason", "crl_reason", "", "revocation reason")
	caCmd.Flags().StringS("crldays", "crldays", "", "Days until the next CRL is due")
	caCmd.Flags().StringS("crlexts", "crlexts", "", "CRL extension section (override value in config file)")
	caCmd.Flags().StringS("crlhours", "crlhours", "", "Hours until the next CRL is due")
	caCmd.Flags().StringS("crlsec", "crlsec", "", "Seconds until the next CRL is due")
	caCmd.Flags().StringSliceS("dateopt", "dateopt", nil, "Datetime format used for printing. (rfc_822/iso_8601). Default is rfc_822.")
	caCmd.Flags().StringS("days", "days", "", "Number of days from today to certify the cert for")
	caCmd.Flags().StringS("enddate", "enddate", "", "[CC]YYMMDDHHMMSSZ value for notAfter certificate field, overrides -days")
	caCmd.Flags().StringS("extensions", "extensions", "", "Extension section (override value in config file)")
	caCmd.Flags().StringS("extfile", "extfile", "", "Configuration file with X509v3 extensions to add")
	caCmd.Flags().BoolS("gencrl", "gencrl", false, "Generate a new CRL")
	caCmd.Flags().StringS("in", "in", "", "The input cert request(s)")
	caCmd.Flags().BoolS("infiles", "infiles", false, "The last argument, requests to process")
	caCmd.Flags().StringS("inform", "inform", "", "CSR input format to use (PEM or DER; by default try PEM first)")
	caCmd.Flags().StringS("key", "key", "", "Key to decrypt the private key or cert files if encrypted. Better use -passin")
	caCmd.Flags().StringS("keyfile", "keyfile", "", "The CA private key")
	caCmd.Flags().StringS("keyform", "keyform", "", "Private key file format (DER/PEM)")
	caCmd.Flags().StringS("md", "md", "", "Digest to use, such as sha256")
	caCmd.Flags().BoolS("msie_hack", "msie_hack", false, "msie modifications to handle all Universal Strings")
	caCmd.Flags().BoolS("multivalue-rdn", "multivalue-rdn", false, "Deprecated; multi-valued RDNs support is always on.")
	caCmd.Flags().StringS("name", "name", "", "The particular CA definition to use")
	caCmd.Flags().BoolS("noemailDN", "noemailDN", false, "Don't add the EMAIL field to the DN")
	caCmd.Flags().StringS("not_after", "not_after", "", "An alias for -enddate")
	caCmd.Flags().StringS("not_before", "not_before", "", "An alias for -startdate")
	caCmd.Flags().BoolS("notext", "notext", false, "Do not print the generated certificate")
	caCmd.Flags().StringS("out", "out", "", "Where to put the output file(s)")
	caCmd.Flags().StringS("outdir", "outdir", "", "Where to put output cert")
	caCmd.Flags().StringS("passin", "passin", "", "Key and cert input file pass phrase source")
	caCmd.Flags().StringS("policy", "policy", "", "The CA 'policy' to support")
	caCmd.Flags().BoolS("preserveDN", "preserveDN", false, "Don't re-order the DN")
	caCmd.Flags().BoolS("quiet", "quiet", false, "Terse output during processing")
	caCmd.Flags().BoolS("rand_serial", "rand_serial", false, "Always create a random serial; do not store it")
	caCmd.Flags().StringS("revoke", "revoke", "", "Revoke a cert (given in file)")
	caCmd.Flags().StringS("section", "section", "", "An alias for -name")
	caCmd.Flags().BoolS("selfsign", "selfsign", false, "Sign a cert with the key associated with it")
	caCmd.Flags().StringSliceS("sigopt", "sigopt", nil, "Signature parameter in n:v form")
	caCmd.Flags().StringS("spkac", "spkac", "", "File contains DN and signed public key and challenge")
	caCmd.Flags().StringS("ss_cert", "ss_cert", "", "File contains a self signed cert to sign")
	caCmd.Flags().StringS("startdate", "startdate", "", "[CC]YYMMDDHHMMSSZ value for notBefore certificate field")
	caCmd.Flags().StringS("status", "status", "", "Shows cert status given the serial number")
	caCmd.Flags().StringS("subj", "subj", "", "Use arg instead of request's subject")
	caCmd.Flags().BoolS("updatedb", "updatedb", false, "Updates db for expired cert")
	caCmd.Flags().BoolS("utf8", "utf8", false, "Input characters are UTF8; default ASCII")
	caCmd.Flags().StringS("valid", "valid", "", "Add a Valid(not-revoked) DB entry about a cert (given in file)")
	caCmd.Flags().BoolS("verbose", "verbose", false, "Verbose output during processing")
	caCmd.Flags().StringSliceS("vfyopt", "vfyopt", nil, "Verification parameter in n:v form")
	common.AddProviderFlags(caCmd)
	common.AddRandomStateFlags(caCmd)
	rootCmd.AddCommand(caCmd)

	caCmd.MarkFlagsMutuallyExclusive("not_after", "enddate")
	caCmd.MarkFlagsMutuallyExclusive("not_before", "startdate")
	caCmd.MarkFlagsMutuallyExclusive("section", "name")

	carapace.Gen(caCmd).FlagCompletion(carapace.ActionMap{
		"cert":     carapace.ActionFiles(),
		"certform": carapace.ActionValues("DER", "PEM", "P12"),
		"config":   carapace.ActionFiles(),
		"extfile":  carapace.ActionFiles(),
		"in":       carapace.ActionFiles(),
		"inform":   carapace.ActionValues("DER", "PEM"),
		"keyfile":  carapace.ActionFiles(),
		"keyform":  carapace.ActionValues("DER", "PEM", "P12"),
		"out":      carapace.ActionFiles(),
		"outdir":   carapace.ActionDirectories(),
		"revoke":   carapace.ActionFiles(),
		"spkac":    carapace.ActionFiles(),
		"ss_cert":  carapace.ActionFiles(),
		"valid":    carapace.ActionFiles(),
	})
}
