package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddValidationFlags(cmd *cobra.Command) {
	cmd.Flags().BoolS("allow_proxy_certs", "allow_proxy_certs", false, "allow the use of proxy certificates")
	cmd.Flags().StringS("attime", "attime", "", "verification epoch time")
	cmd.Flags().StringS("auth_level", "auth_level", "", "chain authentication security level")
	cmd.Flags().BoolS("check_ss_sig", "check_ss_sig", false, "check root CA self-signatures")
	cmd.Flags().BoolS("crl_check", "crl_check", false, "check leaf certificate revocation")
	cmd.Flags().BoolS("crl_check_all", "crl_check_all", false, "check full chain revocation")
	cmd.Flags().BoolS("explicit_policy", "explicit_policy", false, "set policy variable require-explicit-policy")
	cmd.Flags().BoolS("extended_crl", "extended_crl", false, "enable extended CRL features")
	cmd.Flags().BoolS("ignore_critical", "ignore_critical", false, "permit unhandled critical extensions")
	cmd.Flags().BoolS("inhibit_any", "inhibit_any", false, "set policy variable inhibit-any-policy")
	cmd.Flags().BoolS("inhibit_map", "inhibit_map", false, "set policy variable inhibit-policy-mapping")
	cmd.Flags().BoolS("issuer_checks", "issuer_checks", false, "(deprecated)")
	cmd.Flags().BoolS("no_alt_chains", "no_alt_chains", false, "(deprecated)")
	cmd.Flags().BoolS("no_check_time", "no_check_time", false, "ignore certificate validity time")
	cmd.Flags().BoolS("partial_chain", "partial_chain", false, "accept chains anchored by intermediate trust-store CAs")
	cmd.Flags().StringS("policy", "policy", "", "adds policy to the acceptable policy set")
	cmd.Flags().BoolS("policy_check", "policy_check", false, "perform rfc5280 policy checks")
	cmd.Flags().BoolS("policy_print", "policy_print", false, "print policy processing diagnostics")
	cmd.Flags().StringS("purpose", "purpose", "", "certificate chain purpose")
	cmd.Flags().BoolS("suiteB_128", "suiteB_128", false, "Suite B 128-bit mode allowing 192-bit algorithms")
	cmd.Flags().BoolS("suiteB_128_only", "suiteB_128_only", false, "Suite B 128-bit-only mode")
	cmd.Flags().BoolS("suiteB_192", "suiteB_192", false, "Suite B 192-bit-only mode")
	cmd.Flags().BoolS("trusted_first", "trusted_first", false, "search trust store first (default)")
	cmd.Flags().BoolS("use_deltas", "use_deltas", false, "use delta CRLs")
	cmd.Flags().StringS("verify_depth", "verify_depth", "", "chain depth limit")
	cmd.Flags().StringS("verify_email", "verify_email", "", "expected peer email")
	cmd.Flags().StringS("verify_hostname", "verify_hostname", "", "expected peer hostname")
	cmd.Flags().StringS("verify_ip", "verify_ip", "", "expected peer IP address")
	cmd.Flags().StringS("verify_name", "verify_name", "", "verification policy name")
	cmd.Flags().BoolS("x509_strict", "x509_strict", false, "disable certificate compatibility work-arounds")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"policy": carapace.ActionValues(
			"code_sign",
			"default",
			"pkcs7",
			"smime_sign",
			"ssl_client",
			"ssl_server",
		),
		"purpose": carapace.ActionValuesDescribed(
			"sslclient", "SSL client",
			"sslserver", "SSL server",
			"nssslserver", "Netscape SSL server",
			"smimesign", "S/MIME signing",
			"smimeencrypt", "S/MIME encryption",
			"crlsign", "CRL signing",
			"any", "Any Purpose",
			"ocsphelper", "OCSP helper",
			"timestampsign", "Time Stamp signing",
			"codesign", "Code signing",
		),
	})
}
