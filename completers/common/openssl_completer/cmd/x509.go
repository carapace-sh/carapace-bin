package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var x509Cmd = &cobra.Command{
	Use:     "x509",
	Short:   "X.509 Certificate Data Management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(x509Cmd).Standalone()

	x509Cmd.Flags().StringS("CA", "CA", "", "Use the given CA certificate, conflicts with -key")
	x509Cmd.Flags().BoolS("CAcreateserial", "CAcreateserial", false, "Create CA serial number file if it does not exist")
	x509Cmd.Flags().StringS("CAform", "CAform", "", "CA cert format (PEM/DER/P12); has no effect")
	x509Cmd.Flags().StringS("CAkey", "CAkey", "", "The corresponding CA key; default is -CA arg")
	x509Cmd.Flags().StringS("CAkeyform", "CAkeyform", "", "CA key format (ENGINE, other values ignored)")
	x509Cmd.Flags().StringS("CAserial", "CAserial", "", "File that keeps track of CA-generated serial number")
	x509Cmd.Flags().StringS("addreject", "addreject", "", "Reject certificate for a given purpose")
	x509Cmd.Flags().StringS("addtrust", "addtrust", "", "Trust certificate for a given purpose")
	x509Cmd.Flags().BoolS("alias", "alias", false, "Print certificate alias")
	x509Cmd.Flags().BoolS("badsig", "badsig", false, "Corrupt last byte of certificate signature (for test)")
	x509Cmd.Flags().StringSliceS("certopt", "certopt", nil, "Various certificate text printing options")
	x509Cmd.Flags().StringS("checkemail", "checkemail", "", "Check certificate matches email")
	x509Cmd.Flags().StringS("checkend", "checkend", "", "Check whether cert expires in the next arg seconds Exit 1 (failure) if so, 0 if not")
	x509Cmd.Flags().StringS("checkhost", "checkhost", "", "Check certificate matches host")
	x509Cmd.Flags().StringS("checkip", "checkip", "", "Check certificate matches ipaddr")
	x509Cmd.Flags().BoolS("clrext", "clrext", false, "Do not take over any extensions from the source certificate or request")
	x509Cmd.Flags().BoolS("clrreject", "clrreject", false, "Clears all the prohibited or rejected uses of the certificate")
	x509Cmd.Flags().BoolS("clrtrust", "clrtrust", false, "Clear all trusted purposes")
	x509Cmd.Flags().StringS("copy_extensions", "copy_extensions", "", "copy extensions when converting from CSR to x509 or vice versa")
	x509Cmd.Flags().StringSliceS("dateopt", "dateopt", nil, "Datetime format used for printing. (rfc_822/iso_8601). Default is rfc_822.")
	x509Cmd.Flags().BoolS("dates", "dates", false, "Print both notBefore and notAfter fields")
	x509Cmd.Flags().StringS("days", "days", "", "Number of days until newly generated certificate expires - default 30")
	x509Cmd.Flags().BoolS("email", "email", false, "Print email address(es)")
	x509Cmd.Flags().BoolS("enddate", "enddate", false, "Print the notAfter field")
	x509Cmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	x509Cmd.Flags().StringS("ext", "ext", "", "Restrict which X.509 extensions to print and/or copy")
	x509Cmd.Flags().StringS("extensions", "extensions", "", "Section of extfile to use - default: unnamed section")
	x509Cmd.Flags().StringS("extfile", "extfile", "", "Config file with X509V3 extensions to add")
	x509Cmd.Flags().BoolS("fingerprint", "fingerprint", false, "Print the certificate fingerprint")
	x509Cmd.Flags().StringS("force_pubkey", "force_pubkey", "", "Key to be placed in new certificate or certificate request")
	x509Cmd.Flags().BoolS("hash", "hash", false, "Synonym for -subject_hash (for backward compat)")
	x509Cmd.Flags().StringS("in", "in", "", "Certificate input, or CSR input file with -req (default stdin)")
	x509Cmd.Flags().StringS("inform", "inform", "", "CSR input format to use (PEM or DER; by default try PEM first)")
	x509Cmd.Flags().BoolS("issuer", "issuer", false, "Print issuer DN")
	x509Cmd.Flags().BoolS("issuer_hash", "issuer_hash", false, "Print issuer hash value")
	x509Cmd.Flags().BoolS("issuer_hash_old", "issuer_hash_old", false, "Print old-style (MD5) issuer hash value")
	x509Cmd.Flags().StringS("key", "key", "", "Key for signing, and to include unless using -force_pubkey")
	x509Cmd.Flags().StringS("keyform", "keyform", "", "Key input format (ENGINE, other values ignored)")
	x509Cmd.Flags().BoolS("modulus", "modulus", false, "Print the RSA key modulus")
	x509Cmd.Flags().BoolS("multi", "multi", false, "Process multiple certificates")
	x509Cmd.Flags().StringSliceS("nameopt", "nameopt", nil, "Certificate subject/issuer name printing options")
	x509Cmd.Flags().BoolS("new", "new", false, "Generate a certificate from scratch")
	x509Cmd.Flags().BoolS("next_serial", "next_serial", false, "Increment current certificate serial number")
	x509Cmd.Flags().BoolS("nocert", "nocert", false, "No cert output (except for requested printing)")
	x509Cmd.Flags().BoolS("noout", "noout", false, "No output (except for requested printing)")
	x509Cmd.Flags().StringS("not_after", "not_after", "", "[CC]YYMMDDHHMMSSZ value for notAfter certificate field, overrides -days")
	x509Cmd.Flags().StringS("not_before", "not_before", "", "[CC]YYMMDDHHMMSSZ value for notBefore certificate field")
	x509Cmd.Flags().BoolS("ocsp_uri", "ocsp_uri", false, "Print OCSP Responder URL(s)")
	x509Cmd.Flags().BoolS("ocspid", "ocspid", false, "Print OCSP hash values for the subject name and public key")
	x509Cmd.Flags().StringS("out", "out", "", "Output file - default stdout")
	x509Cmd.Flags().StringS("outform", "outform", "", "Output format (DER or PEM) - default PEM")
	x509Cmd.Flags().StringS("passin", "passin", "", "Private key and cert file pass-phrase source")
	x509Cmd.Flags().BoolS("preserve_dates", "preserve_dates", false, "Preserve existing validity dates")
	x509Cmd.Flags().BoolS("pubkey", "pubkey", false, "Print the public key in PEM format")
	x509Cmd.Flags().BoolS("purpose", "purpose", false, "Print out certificate purposes")
	x509Cmd.Flags().BoolS("req", "req", false, "Input is a CSR file (rather than a certificate)")
	x509Cmd.Flags().BoolS("serial", "serial", false, "Print serial number value")
	x509Cmd.Flags().StringS("set_issuer", "set_issuer", "", "Set or override certificate issuer")
	x509Cmd.Flags().StringS("set_serial", "set_serial", "", "Serial number to use, overrides -CAserial")
	x509Cmd.Flags().StringS("set_subject", "set_subject", "", "Set or override certificate subject (and issuer)")
	x509Cmd.Flags().StringS("setalias", "setalias", "", "Set certificate alias (nickname)")
	x509Cmd.Flags().StringS("signkey", "signkey", "", "Same as -key")
	x509Cmd.Flags().StringSliceS("sigopt", "sigopt", nil, "Signature parameter, in n:v form")
	x509Cmd.Flags().BoolS("startdate", "startdate", false, "Print the notBefore field")
	x509Cmd.Flags().StringS("subj", "subj", "", "Alias for -set_subject")
	x509Cmd.Flags().BoolS("subject", "subject", false, "Print subject DN")
	x509Cmd.Flags().BoolS("subject_hash", "subject_hash", false, "Print subject hash value")
	x509Cmd.Flags().BoolS("subject_hash_old", "subject_hash_old", false, "Print old-style (MD5) subject hash value")
	x509Cmd.Flags().BoolS("text", "text", false, "Print the certificate in text form")
	x509Cmd.Flags().BoolS("trustout", "trustout", false, "Mark certificate PEM output as trusted")
	x509Cmd.Flags().StringSliceS("vfyopt", "vfyopt", nil, "CSR verification parameter in n:v form")
	x509Cmd.Flags().BoolS("x509toreq", "x509toreq", false, "Output a certification request (rather than a certificate)")
	rootCmd.AddCommand(x509Cmd)

	carapace.Gen(x509Cmd).FlagCompletion(carapace.ActionMap{
		"CA":           carapace.ActionFiles(),
		"CAform":       carapace.ActionValues("DER", "PEM", "P12"),
		"CAkey":        carapace.ActionFiles(),
		"CAkeyform":    carapace.ActionValues("ENGINE", "DER", "PEM", "P12"),
		"CAserial":     carapace.ActionFiles(),
		"extfile":      carapace.ActionFiles(),
		"force_pubkey": carapace.ActionFiles(),
		"in":           carapace.ActionFiles(),
		"inform":       carapace.ActionValues("DER", "PEM"),
		"key":          carapace.ActionFiles(),
		"keyform":      carapace.ActionValues("ENGINE", "DER", "PEM", "P12"),
		"out":          carapace.ActionFiles(),
		"outform":      carapace.ActionValues("DER", "PEM"),
		"signkey":      carapace.ActionFiles(),
	})
}
