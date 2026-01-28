package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var cmpCmd = &cobra.Command{
	Use:     "cmp",
	Short:   "Certificate Management Protocol (CMP, RFCs 9810 and 9811) application",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cmpCmd).Standalone()

	cmpCmd.Flags().BoolS("accept_raverified", "accept_raverified", false, "Accept RAVERIFIED as proof-of-possession (POPO)")
	cmpCmd.Flags().BoolS("accept_unprot_err", "accept_unprot_err", false, "Accept unprotected error messages from client")
	cmpCmd.Flags().BoolS("accept_unprotected", "accept_unprotected", false, "Accept missing or invalid protection of requests")
	cmpCmd.Flags().BoolS("batch", "batch", false, "Do not interactively prompt for input when a password is required etc.")
	cmpCmd.Flags().StringS("cacertsout", "cacertsout", "", "File to save CA certs received in caPubs field or genp with id-it-caCerts")
	cmpCmd.Flags().BoolS("centralkeygen", "centralkeygen", false, "Request central (server-side) key generation. Default is local generation")
	cmpCmd.Flags().StringS("cert", "cert", "", "Client's CMP signer certificate; its public key must match the -key argument This also used as default reference for subject DN and SANs. Any further certs included are appended to the untrusted certs")
	cmpCmd.Flags().StringS("certform", "certform", "", "Format (PEM or DER) to use when saving a certificate to a file. Default PEM")
	cmpCmd.Flags().StringS("certout", "certout", "", "File to save newly enrolled certificate")
	cmpCmd.Flags().StringS("chainout", "chainout", "", "File to save the chain of newly enrolled certificate")
	cmpCmd.Flags().StringS("check_after", "check_after", "", "The check_after value (time to wait) to include in poll response")
	cmpCmd.Flags().StringS("cmd", "cmd", "", "CMP request to send: ir/cr/kur/p10cr/rr/genm")
	cmpCmd.Flags().StringS("config", "config", "", "Configuration file to use. \"\" = none. Default from env variable OPENSSL_CONF")
	cmpCmd.Flags().StringS("crlcert", "crlcert", "", "certificate to request a CRL for in genm of type crlStatusList")
	cmpCmd.Flags().StringS("crlform", "crlform", "", "Format (PEM or DER) to use when saving a CRL to a file. Default DER")
	cmpCmd.Flags().StringS("crlout", "crlout", "", "File to save new CRL received in genp of type 'crls'")
	cmpCmd.Flags().StringS("csr", "csr", "", "PKCS#10 CSR file in PEM or DER format to convert or to use in p10cr")
	cmpCmd.Flags().StringS("days", "days", "", "Requested validity time of the new certificate in number of days")
	cmpCmd.Flags().StringS("digest", "digest", "", "Digest to use in message protection and POPO signatures. Default \"sha256\"")
	cmpCmd.Flags().BoolS("disable_confirm", "disable_confirm", false, "Do not confirm newly enrolled certificate w/o requesting implicit confirmation. WARNING: This leads to behavior violating RFC 9810")
	cmpCmd.Flags().StringS("expect_sender", "expect_sender", "", "DN of expected sender of responses. Defaults to subject of -srvcert, if any")
	cmpCmd.Flags().StringS("extracerts", "extracerts", "", "Certificates to append in extraCerts field of outgoing messages. This can be used as the default CMP signer cert chain to include")
	cmpCmd.Flags().StringS("extracertsout", "extracertsout", "", "File to save extra certificates received in the extraCerts field")
	cmpCmd.Flags().StringS("failure", "failure", "", "A single failure info bit number to include in server response, 0..26")
	cmpCmd.Flags().StringS("failurebits", "failurebits", "", "Number representing failure bits to include in server response, 0..2^27 - 1")
	cmpCmd.Flags().StringS("geninfo", "geninfo", "", "Comma-separated list of OID and value to place in generalInfo PKIHeader of form <OID>:int:<n> or <OID>:str:<s>, e.g. '1.2.3.4:int:56789, id-kp:str:name'")
	cmpCmd.Flags().BoolS("grant_implicitconf", "grant_implicitconf", false, "Grant implicit confirmation of newly enrolled certificate")
	cmpCmd.Flags().BoolS("ignore_keyusage", "ignore_keyusage", false, "Ignore CMP signer cert key usage, else 'digitalSignature' must be allowed")
	cmpCmd.Flags().BoolS("implicit_confirm", "implicit_confirm", false, "Request implicit confirmation of newly enrolled certificates")
	cmpCmd.Flags().StringS("infotype", "infotype", "", "InfoType name for requesting specific info in genm, with specific support for 'caCerts' and 'rootCaCert'")
	cmpCmd.Flags().StringS("issuer", "issuer", "", "DN of the issuer to place in the certificate template of ir/cr/kur/rr; also used as recipient if neither -recipient nor -srvcert are given")
	cmpCmd.Flags().StringS("keep_alive", "keep_alive", "", "Persistent HTTP connections. 0: no, 1 (the default): request, 2: require")
	cmpCmd.Flags().StringS("key", "key", "", "CMP signer private key, not used when -secret given")
	cmpCmd.Flags().StringS("keyform", "keyform", "", "Format of the key input (DER/PEM/P12)")
	cmpCmd.Flags().StringS("keypass", "keypass", "", "Client private key (and cert and old cert) pass phrase source")
	cmpCmd.Flags().StringS("keyspec", "keyspec", "", "Optional file to save Key specification received in genp of type certReqTemplate")
	cmpCmd.Flags().StringS("mac", "mac", "", "MAC algorithm to use in PBM-based message protection. Default \"hmac-sha1\"")
	cmpCmd.Flags().StringS("max_msgs", "max_msgs", "", "max number of messages handled by HTTP mock server. Default: 0 = unlimited")
	cmpCmd.Flags().StringS("msg_timeout", "msg_timeout", "", "Number of seconds allowed per CMP message round trip, or 0 for infinite")
	cmpCmd.Flags().StringS("newkey", "newkey", "", "Private or public key for the requested cert. Default: CSR key or client key")
	cmpCmd.Flags().StringS("newkeyout", "newkeyout", "", "File to save centrally generated key, in PEM format")
	cmpCmd.Flags().StringS("newkeypass", "newkeypass", "", "New private key pass phrase source")
	cmpCmd.Flags().StringS("newwithnew", "newwithnew", "", "File to save NewWithNew cert received in genp of type rootCaKeyUpdate")
	cmpCmd.Flags().StringS("newwithold", "newwithold", "", "File to save NewWithOld cert received in genp of type rootCaKeyUpdate")
	cmpCmd.Flags().BoolS("no_cache_extracerts", "no_cache_extracerts", false, "Do not keep certificates received in the extraCerts CMP message field")
	cmpCmd.Flags().StringS("no_proxy", "no_proxy", "", "List of addresses of servers not to use HTTP(S) proxy for Default from environment variable 'no_proxy', else 'NO_PROXY', else none")
	cmpCmd.Flags().StringS("oldcert", "oldcert", "", "Certificate to be updated (defaulting to -cert) or to be revoked in rr; also used as reference (defaulting to -cert) for subject DN and SANs. Issuer is used as recipient unless -recipient, -srvcert, or -issuer given")
	cmpCmd.Flags().StringS("oldcrl", "oldcrl", "", "CRL to request update for in genm of type crlStatusList")
	cmpCmd.Flags().StringS("oldwithnew", "oldwithnew", "", "File to save OldWithNew cert received in genp of type rootCaKeyUpdate")
	cmpCmd.Flags().StringS("oldwithold", "oldwithold", "", "Root CA certificate to request update for in genm of type rootCaCert")
	cmpCmd.Flags().StringS("otherpass", "otherpass", "", "Pass phrase source potentially needed for loading certificates of others")
	cmpCmd.Flags().StringS("out_trusted", "out_trusted", "", "Certificates to trust when verifying newly enrolled certificates")
	cmpCmd.Flags().StringS("own_trusted", "own_trusted", "", "Optional certs to verify chain building for own CMP signer cert")
	cmpCmd.Flags().StringS("path", "path", "", "HTTP path (aka CMP alias) at the CMP server. Default from -server, else \"/\"")
	cmpCmd.Flags().StringS("pkistatus", "pkistatus", "", "PKIStatus to be included in server response. Possible values: 0..6")
	cmpCmd.Flags().StringS("policies", "policies", "", "Name of config file section defining policies certificate request extension")
	cmpCmd.Flags().StringS("policy_oids", "policy_oids", "", "Policy OID(s) to add as policies certificate request extension")
	cmpCmd.Flags().BoolS("policy_oids_critical", "policy_oids_critical", false, "Flag the policy OID(s) given with -policy_oids as critical")
	cmpCmd.Flags().StringS("poll_count", "poll_count", "", "Number of times the client must poll before receiving a certificate")
	cmpCmd.Flags().StringS("popo", "popo", "", "Proof-of-Possession (POPO) method to use for ir/cr/kur where -1 = NONE, 0 = RAVERIFIED, 1 = SIGNATURE (default), 2 = KEYENC")
	cmpCmd.Flags().StringS("port", "port", "", "Act as HTTP-based mock server listening on given port")
	cmpCmd.Flags().StringS("profile", "profile", "", "Certificate profile name to place in generalInfo field of request PKIHeader")
	cmpCmd.Flags().StringS("proxy", "proxy", "", "[http[s]://]address[:port][/path] of HTTP(S) proxy to use; path is ignored")
	cmpCmd.Flags().StringS("recipient", "recipient", "", "DN of CA. Default: subject of -srvcert, -issuer, issuer of -oldcert or -cert")
	cmpCmd.Flags().StringS("ref", "ref", "", "Reference value to use as senderKID in case no -cert is given")
	cmpCmd.Flags().StringS("ref_cert", "ref_cert", "", "Certificate to be expected for rr and any oldCertID in kur messages")
	cmpCmd.Flags().StringS("repeat", "repeat", "", "Invoke the transaction the given positive number of times. Default 1")
	cmpCmd.Flags().StringS("reqexts", "reqexts", "", "Name of config file section defining certificate request extensions. Augments or replaces any extensions contained CSR given with -csr")
	cmpCmd.Flags().StringS("reqin", "reqin", "", "Take sequence of CMP requests to send to server from file(s)")
	cmpCmd.Flags().BoolS("reqin_new_tid", "reqin_new_tid", false, "Use fresh transactionID for CMP requests read from -reqin")
	cmpCmd.Flags().StringS("reqout", "reqout", "", "Save sequence of CMP requests created by the client to file(s)")
	cmpCmd.Flags().StringS("reqout_only", "reqout_only", "", "Save first CMP request created by the client to file and exit")
	cmpCmd.Flags().StringS("revreason", "revreason", "", "Reason code to include in revocation request (rr); possible values: 0..6, 8..10 (see RFC5280, 5.3.1) or -1. Default -1 = none included")
	cmpCmd.Flags().StringS("rsp_capubs", "rsp_capubs", "", "CA certificates to be included in mock ip response")
	cmpCmd.Flags().StringS("rsp_cert", "rsp_cert", "", "Certificate to be returned as mock enrollment result")
	cmpCmd.Flags().StringS("rsp_crl", "rsp_crl", "", "CRL to be returned in genp of type crls")
	cmpCmd.Flags().StringS("rsp_extracerts", "rsp_extracerts", "", "Extra certificates to be included in mock certification responses")
	cmpCmd.Flags().StringS("rsp_key", "rsp_key", "", "Private key for the certificate to be returned as mock enrollment result Key to be returned for central key pair generation")
	cmpCmd.Flags().StringS("rsp_keypass", "rsp_keypass", "", "Response private key (and cert) pass phrase source")
	cmpCmd.Flags().StringS("rsp_newwithnew", "rsp_newwithnew", "", "New root CA certificate to include in genp of type rootCaKeyUpdate")
	cmpCmd.Flags().StringS("rsp_newwithold", "rsp_newwithold", "", "NewWithOld transition cert to include in genp of type rootCaKeyUpdate")
	cmpCmd.Flags().StringS("rsp_oldwithnew", "rsp_oldwithnew", "", "OldWithNew transition cert to include in genp of type rootCaKeyUpdate")
	cmpCmd.Flags().StringS("rspin", "rspin", "", "Process sequence of CMP responses provided in file(s), skipping server")
	cmpCmd.Flags().StringS("rspout", "rspout", "", "Save sequence of actually used CMP responses to file(s)")
	cmpCmd.Flags().BoolS("san_nodefault", "san_nodefault", false, "Do not take default SANs from reference certificate (see -oldcert)")
	cmpCmd.Flags().StringS("sans", "sans", "", "Subject Alt Names (IPADDR/DNS/URI) to add as (critical) cert req extension")
	cmpCmd.Flags().StringS("secret", "secret", "", "Prefer PBM (over signatures) for protecting msgs with given password source")
	cmpCmd.Flags().StringS("section", "section", "", "Section(s) in config file to get options from. \"\" = 'default'. Default 'cmp'")
	cmpCmd.Flags().BoolS("send_error", "send_error", false, "Force server to reply with error message")
	cmpCmd.Flags().BoolS("send_unprot_err", "send_unprot_err", false, "In case of negative responses, server shall send unprotected error messages, certificate responses (ip/cp/kup), and revocation responses (rp). WARNING: This setting leads to behavior violating RFC 9810")
	cmpCmd.Flags().BoolS("send_unprotected", "send_unprotected", false, "Send response messages without CMP-level protection")
	cmpCmd.Flags().StringS("serial", "serial", "", "Serial number of certificate to be revoked in revocation request (rr)")
	cmpCmd.Flags().StringS("server", "server", "", "[http[s]://]address[:port][/path] of CMP server. Default port 80 or 443. address may be a DNS name or an IP address; path can be overridden by -path")
	cmpCmd.Flags().StringS("srv_cert", "srv_cert", "", "Certificate of the server")
	cmpCmd.Flags().StringS("srv_key", "srv_key", "", "Private key used by the server for signing messages")
	cmpCmd.Flags().StringS("srv_keypass", "srv_keypass", "", "Server private key (and cert) pass phrase source")
	cmpCmd.Flags().StringS("srv_ref", "srv_ref", "", "Reference value to use as senderKID of server in case no -srv_cert is given")
	cmpCmd.Flags().StringS("srv_secret", "srv_secret", "", "Password source for server authentication with a pre-shared key (secret)")
	cmpCmd.Flags().StringS("srv_trusted", "srv_trusted", "", "Trusted certificates for client authentication")
	cmpCmd.Flags().StringS("srv_untrusted", "srv_untrusted", "", "Intermediate certs that may be useful for verifying CMP protection")
	cmpCmd.Flags().StringS("srvcert", "srvcert", "", "Server cert to pin and trust directly when verifying signed CMP responses")
	cmpCmd.Flags().StringS("srvcertout", "srvcertout", "", "File to save the server cert used and validated for CMP response protection")
	cmpCmd.Flags().StringS("statusstring", "statusstring", "", "Status string to be included in server response")
	cmpCmd.Flags().StringS("subject", "subject", "", "Distinguished Name (DN) of subject to use in the requested cert template For kur, default is subject of -csr arg or reference cert (see -oldcert) this default is used for ir and cr only if no Subject Alt Names are set")
	cmpCmd.Flags().BoolS("ta_in_ip_extracerts", "ta_in_ip_extracerts", false, "Permit using self-issued certificates from the extraCerts in an IP message as trust anchors under conditions defined by 3GPP TS 33.310 WARNING: This setting leads to behavior allowing violation of RFC 9810")
	cmpCmd.Flags().StringS("template", "template", "", "File to save certTemplate received in genp of type certReqTemplate")
	cmpCmd.Flags().StringS("tls_cert", "tls_cert", "", "Client's TLS certificate. May include chain to be provided to TLS server")
	cmpCmd.Flags().StringS("tls_extra", "tls_extra", "", "Extra certificates to provide to TLS server during TLS handshake")
	cmpCmd.Flags().StringS("tls_host", "tls_host", "", "Address to be checked (rather than -server) during TLS hostname validation")
	cmpCmd.Flags().StringS("tls_key", "tls_key", "", "Private key for the client's TLS certificate")
	cmpCmd.Flags().StringS("tls_keypass", "tls_keypass", "", "Pass phrase source for the client's private TLS key (and TLS cert)")
	cmpCmd.Flags().StringS("tls_trusted", "tls_trusted", "", "Trusted certificates to use for verifying the TLS server certificate; this implies hostname validation")
	cmpCmd.Flags().BoolS("tls_used", "tls_used", false, "Enable using TLS (also when other TLS options are not set)")
	cmpCmd.Flags().StringS("total_timeout", "total_timeout", "", "Overall time an enrollment incl. polling may take. Default 0 = infinite")
	cmpCmd.Flags().StringS("trusted", "trusted", "", "Certificates to use as trust anchors when verifying signed CMP responses unless -srvcert is given")
	cmpCmd.Flags().BoolS("unprotected_errors", "unprotected_errors", false, "Accept missing or invalid protection of regular error messages and negative certificate responses (ip/cp/kup), revocation responses (rp), and PKIConf WARNING: This setting leads to behavior allowing violation of RFC 9810")
	cmpCmd.Flags().BoolS("unprotected_requests", "unprotected_requests", false, "Send request messages without CMP-level protection")
	cmpCmd.Flags().StringS("untrusted", "untrusted", "", "Intermediate CA certs for chain construction for CMP/TLS/enrolled certs")
	cmpCmd.Flags().BoolS("use_mock_srv", "use_mock_srv", false, "Use internal mock server at API level, bypassing socket-based HTTP")
	cmpCmd.Flags().StringS("verbosity", "verbosity", "", "Log level; 3=ERR, 4=WARN, 6=INFO, 7=DEBUG, 8=TRACE. Default 6 = INFO")
	common.AddProviderFlags(cmpCmd)
	common.AddRandomStateFlags(cmpCmd)
	common.AddValidationFlags(cmpCmd)
	rootCmd.AddCommand(cmpCmd)

	carapace.Gen(cmpCmd).FlagCompletion(carapace.ActionMap{
		"cacertsout": carapace.ActionFiles(),
		"cert":       carapace.ActionFiles(),
		"certform":   carapace.ActionValues("DER", "PEM"),
		"certout":    carapace.ActionFiles(),
		"chainout":   carapace.ActionFiles(),
		"cmd": carapace.ActionValuesDescribed(
			"ir", "Initialization Request",
			"cr", "Certificate Request",
			"p10cr", "PKCS#10 Certification Request (for legacy support)",
			"kur", "Key Update Request",
			"rr", "Revocation Request",
			"genm", "General Message",
		),
		"config":         carapace.ActionFiles(),
		"crlcert":        carapace.ActionFiles(),
		"crlform":        carapace.ActionValues("DER", "PEM"),
		"crlout":         carapace.ActionFiles(),
		"csr":            carapace.ActionFiles(),
		"extracerts":     carapace.ActionFiles(),
		"extracertsout":  carapace.ActionFiles(),
		"key":            carapace.ActionFiles(),
		"keyform":        carapace.ActionValues("DER", "PEM", "P12"),
		"keyspec":        carapace.ActionFiles(),
		"newkey":         carapace.ActionFiles(),
		"newkeyout":      carapace.ActionFiles(),
		"newwithnew":     carapace.ActionFiles(),
		"newwithold":     carapace.ActionFiles(),
		"oldcert":        carapace.ActionFiles(),
		"oldcrl":         carapace.ActionFiles(),
		"oldwithnew":     carapace.ActionFiles(),
		"oldwithold":     carapace.ActionFiles(),
		"out_trusted":    carapace.ActionFiles(),
		"own_trusted":    carapace.ActionFiles(),
		"ref_cert":       carapace.ActionFiles(),
		"reqin":          carapace.ActionFiles(),
		"reqout":         carapace.ActionFiles(),
		"reqout_only":    carapace.ActionFiles(),
		"rsp_capubs":     carapace.ActionFiles(),
		"rsp_cert":       carapace.ActionFiles(),
		"rsp_crl":        carapace.ActionFiles(),
		"rsp_extracerts": carapace.ActionFiles(),
		"rsp_key":        carapace.ActionFiles(),
		"rsp_newwithnew": carapace.ActionFiles(),
		"rsp_newwithold": carapace.ActionFiles(),
		"rspin":          carapace.ActionFiles(),
		"rspout":         carapace.ActionFiles(),
		"srv_cert":       carapace.ActionFiles(),
		"srv_key":        carapace.ActionFiles(),
		"srv_trusted":    carapace.ActionFiles(),
		"srv_untrusted":  carapace.ActionFiles(),
		"srvcert":        carapace.ActionFiles(),
		"srvcertout":     carapace.ActionFiles(),
		"template":       carapace.ActionFiles(),
		"tls_cert":       carapace.ActionFiles(),
		"tls_extra":      carapace.ActionFiles(),
		"tls_key":        carapace.ActionFiles(),
		"tls_trusted":    carapace.ActionFiles(),
		"trusted":        carapace.ActionFiles(),
		"untrusted":      carapace.ActionFiles(),
	})
}
