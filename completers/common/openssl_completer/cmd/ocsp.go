package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var ocspCmd = &cobra.Command{
	Use:     "ocsp",
	Short:   "Online Certificate Status Protocol command",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ocspCmd).Standalone()

	ocspCmd.Flags().StringS("CA", "CA", "", "CA certificates")
	ocspCmd.Flags().StringS("CAfile", "CAfile", "", "Trusted certificates file")
	ocspCmd.Flags().StringS("CApath", "CApath", "", "Trusted certificates directory")
	ocspCmd.Flags().StringS("CAstore", "CAstore", "", "Trusted certificates store URI")
	ocspCmd.Flags().StringS("VAfile", "VAfile", "", "Validator certificates file")
	ocspCmd.Flags().BoolS("badsig", "badsig", false, "Corrupt last byte of loaded OCSP response signature (for test)")
	ocspCmd.Flags().StringSliceS("cert", "cert", nil, "Certificate to check; may be given multiple times")
	ocspCmd.Flags().StringS("header", "header", "", "key=value header to add")
	ocspCmd.Flags().StringS("host", "host", "", "TCP/IP hostname:port to connect to")
	ocspCmd.Flags().BoolS("ignore_err", "ignore_err", false, "Ignore error on OCSP request or response and continue running")
	ocspCmd.Flags().StringS("index", "index", "", "Certificate status index file")
	ocspCmd.Flags().StringS("issuer", "issuer", "", "Issuer certificate")
	ocspCmd.Flags().StringS("multi", "multi", "", "run multiple responder processes")
	ocspCmd.Flags().StringS("ndays", "ndays", "", "Number of days before next update")
	ocspCmd.Flags().StringS("nmin", "nmin", "", "Number of minutes before next update")
	ocspCmd.Flags().BoolS("no-CAfile", "no-CAfile", false, "Do not load the default certificates file")
	ocspCmd.Flags().BoolS("no-CApath", "no-CApath", false, "Do not load certificates from the default certificates directory")
	ocspCmd.Flags().BoolS("no-CAstore", "no-CAstore", false, "Do not load certificates from the default certificates store")
	ocspCmd.Flags().BoolS("no_cert_checks", "no_cert_checks", false, "Don't do additional checks on signing certificate")
	ocspCmd.Flags().BoolS("no_cert_verify", "no_cert_verify", false, "Don't check signing certificate")
	ocspCmd.Flags().BoolS("no_certs", "no_certs", false, "Don't include any certificates in signed request")
	ocspCmd.Flags().BoolS("no_chain", "no_chain", false, "Don't chain verify response")
	ocspCmd.Flags().BoolS("no_explicit", "no_explicit", false, "Do not explicitly check the chain, just verify the root")
	ocspCmd.Flags().BoolS("no_intern", "no_intern", false, "Don't search certificates contained in response for signer")
	ocspCmd.Flags().BoolS("no_nonce", "no_nonce", false, "Don't add OCSP nonce to request")
	ocspCmd.Flags().StringS("no_proxy", "no_proxy", "", "List of addresses of servers not to use HTTP(S) proxy for Default from environment variable 'no_proxy', else 'NO_PROXY', else none")
	ocspCmd.Flags().BoolS("no_signature_verify", "no_signature_verify", false, "Don't check signature on response")
	ocspCmd.Flags().BoolS("nonce", "nonce", false, "Add OCSP nonce to request")
	ocspCmd.Flags().BoolS("noverify", "noverify", false, "Don't verify response at all")
	ocspCmd.Flags().StringS("nrequest", "nrequest", "", "Number of requests to accept (default unlimited)")
	ocspCmd.Flags().StringS("out", "out", "", "Output filename")
	ocspCmd.Flags().StringS("passin", "passin", "", "Responder key pass phrase source")
	ocspCmd.Flags().StringS("path", "path", "", "Path to use in OCSP request")
	ocspCmd.Flags().StringS("port", "port", "", "Port to run responder on")
	ocspCmd.Flags().StringS("proxy", "proxy", "", "[http[s]://]host[:port][/path] of HTTP(S) proxy to use; path is ignored")
	ocspCmd.Flags().StringS("rcid", "rcid", "", "Use specified algorithm for cert id in response")
	ocspCmd.Flags().BoolS("req_text", "req_text", false, "Print text form of request")
	ocspCmd.Flags().StringS("reqin", "reqin", "", "File with the DER-encoded request")
	ocspCmd.Flags().StringS("reqout", "reqout", "", "Output file for the DER-encoded request")
	ocspCmd.Flags().BoolS("resp_key_id", "resp_key_id", false, "Identify response by signing certificate key ID")
	ocspCmd.Flags().BoolS("resp_no_certs", "resp_no_certs", false, "Don't include any certificates in response")
	ocspCmd.Flags().BoolS("resp_text", "resp_text", false, "Print text form of response")
	ocspCmd.Flags().StringS("respin", "respin", "", "File with the DER-encoded response")
	ocspCmd.Flags().StringS("respout", "respout", "", "Output file for the DER-encoded response")
	ocspCmd.Flags().StringS("rkey", "rkey", "", "Responder key to sign responses with")
	ocspCmd.Flags().StringS("rmd", "rmd", "", "Digest Algorithm to use in signature of OCSP response")
	ocspCmd.Flags().StringS("rother", "rother", "", "Other certificates to include in response")
	ocspCmd.Flags().StringS("rsigner", "rsigner", "", "Responder certificate to sign responses with")
	ocspCmd.Flags().StringSliceS("rsigopt", "rsigopt", nil, "OCSP response signature parameter in n:v form")
	ocspCmd.Flags().StringS("serial", "serial", "", "Serial number to check; may be given multiple times")
	ocspCmd.Flags().StringS("sign_other", "sign_other", "", "Additional certificates to include in signed request")
	ocspCmd.Flags().StringS("signer", "signer", "", "Certificate to sign OCSP request with")
	ocspCmd.Flags().StringS("signkey", "signkey", "", "Private key to sign OCSP request with")
	ocspCmd.Flags().StringS("status_age", "status_age", "", "Maximum status age in seconds")
	ocspCmd.Flags().BoolS("text", "text", false, "Print text form of request and response")
	ocspCmd.Flags().StringS("timeout", "timeout", "", "Connection timeout (in seconds) to the OCSP responder")
	ocspCmd.Flags().BoolS("trust_other", "trust_other", false, "Don't verify additional certificates")
	ocspCmd.Flags().StringS("url", "url", "", "Responder URL")
	ocspCmd.Flags().StringS("validity_period", "validity_period", "", "Maximum validity discrepancy in seconds")
	ocspCmd.Flags().StringS("verify_other", "verify_other", "", "Additional certificates to search for signer")
	common.AddProviderFlags(ocspCmd)
	common.AddValidationFlags(ocspCmd)
	rootCmd.AddCommand(ocspCmd)

	carapace.Gen(ocspCmd).FlagCompletion(carapace.ActionMap{
		"CA":           carapace.ActionFiles(),
		"CAfile":       carapace.ActionFiles(),
		"CApath":       carapace.ActionDirectories(),
		"VAfile":       carapace.ActionFiles(),
		"cert":         carapace.ActionFiles(),
		"index":        carapace.ActionFiles(),
		"issuer":       carapace.ActionFiles(),
		"out":          carapace.ActionFiles(),
		"reqin":        carapace.ActionFiles(),
		"reqout":       carapace.ActionFiles(),
		"respin":       carapace.ActionFiles(),
		"respout":      carapace.ActionFiles(),
		"rkey":         carapace.ActionFiles(),
		"rother":       carapace.ActionFiles(),
		"rsigner":      carapace.ActionFiles(),
		"sign_other":   carapace.ActionFiles(),
		"signer":       carapace.ActionFiles(),
		"signkey":      carapace.ActionFiles(),
		"verify_other": carapace.ActionFiles(),
	})
}
