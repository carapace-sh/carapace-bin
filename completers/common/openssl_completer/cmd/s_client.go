package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var s_clientCmd = &cobra.Command{
	Use:     "s_client",
	Short:   "SSL/TLS client program",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s_clientCmd).Standalone()

	s_clientCmd.Flags().BoolS("4", "4", false, "Use IPv4 only")
	s_clientCmd.Flags().BoolS("6", "6", false, "Use IPv6 only")
	s_clientCmd.Flags().StringS("CAfile", "CAfile", "", "PEM format file of CA's")
	s_clientCmd.Flags().StringS("CApath", "CApath", "", "PEM format directory of CA's")
	s_clientCmd.Flags().StringS("CAstore", "CAstore", "", "URI to store of CA's")
	s_clientCmd.Flags().StringS("CRL", "CRL", "", "CRL file to use")
	s_clientCmd.Flags().StringS("CRLform", "CRLform", "", "CRL format (PEM or DER); default PEM")
	s_clientCmd.Flags().BoolS("adv", "adv", false, "Advanced command mode")
	s_clientCmd.Flags().StringS("alpn", "alpn", "", "Enable ALPN extension, considering named protocols supported (comma-separated list)")
	s_clientCmd.Flags().BoolS("async", "async", false, "Support asynchronous operation")
	s_clientCmd.Flags().StringS("bind", "bind", "", "bind local address for connection")
	s_clientCmd.Flags().BoolS("brief", "brief", false, "Restrict output to brief summary of connection parameters")
	s_clientCmd.Flags().BoolS("build_chain", "build_chain", false, "Build client certificate chain")
	s_clientCmd.Flags().StringS("cert", "cert", "", "Client certificate file to use")
	s_clientCmd.Flags().StringS("cert_chain", "cert_chain", "", "Client certificate chain file (in PEM format)")
	s_clientCmd.Flags().StringS("certform", "certform", "", "Client certificate file format (PEM/DER/P12); has no effect")
	s_clientCmd.Flags().StringS("chainCAfile", "chainCAfile", "", "CA file for certificate chain (PEM format)")
	s_clientCmd.Flags().StringS("chainCApath", "chainCApath", "", "Use dir as certificate store path to build CA certificate chain")
	s_clientCmd.Flags().StringS("chainCAstore", "chainCAstore", "", "CA store URI for certificate chain")
	s_clientCmd.Flags().StringS("connect", "connect", "", "TCP/IP where to connect; default: 4433)")
	s_clientCmd.Flags().BoolS("crl_download", "crl_download", false, "Download CRL from distribution points")
	s_clientCmd.Flags().BoolS("crlf", "crlf", false, "Convert LF from terminal into CRLF")
	s_clientCmd.Flags().BoolS("ct", "ct", false, "Request and parse SCTs (also enables OCSP stapling)")
	s_clientCmd.Flags().StringS("ctlogfile", "ctlogfile", "", "CT log list CONF file")
	s_clientCmd.Flags().BoolS("dane_ee_no_namechecks", "dane_ee_no_namechecks", false, "Disable name checks when matching DANE-EE(3) TLSA records")
	s_clientCmd.Flags().StringS("dane_tlsa_domain", "dane_tlsa_domain", "", "DANE TLSA base domain")
	s_clientCmd.Flags().StringS("dane_tlsa_rrdata", "dane_tlsa_rrdata", "", "DANE TLSA rrdata presentation form")
	s_clientCmd.Flags().BoolS("debug", "debug", false, "Extra output")
	s_clientCmd.Flags().BoolS("dtls", "dtls", false, "Use any version of DTLS")
	s_clientCmd.Flags().BoolS("dtls1", "dtls1", false, "Just use DTLSv1")
	s_clientCmd.Flags().BoolS("dtls1_2", "dtls1_2", false, "Just use DTLSv1.2")
	s_clientCmd.Flags().StringS("early_data", "early_data", "", "File to send as early data")
	s_clientCmd.Flags().BoolS("enable_client_rpk", "enable_client_rpk", false, "Enable raw public keys (RFC7250) from the client")
	s_clientCmd.Flags().BoolS("enable_pha", "enable_pha", false, "Enable post-handshake-authentication")
	s_clientCmd.Flags().BoolS("enable_server_rpk", "enable_server_rpk", false, "Enable raw public keys (RFC7250) from the server")
	s_clientCmd.Flags().BoolS("fallback_scsv", "fallback_scsv", false, "Send the fallback SCSV")
	s_clientCmd.Flags().StringS("host", "host", "", "Use -connect instead")
	s_clientCmd.Flags().BoolS("ign_eof", "ign_eof", false, "Ignore input eof (default when -quiet)")
	s_clientCmd.Flags().BoolS("ignore_unexpected_eof", "ignore_unexpected_eof", false, "Do not treat lack of close_notify from a peer as an error")
	s_clientCmd.Flags().StringS("key", "key", "", "Private key file to use; default: -cert file")
	s_clientCmd.Flags().StringS("keyform", "keyform", "", "Key format (DER/PEM)")
	s_clientCmd.Flags().StringS("keylogfile", "keylogfile", "", "Write TLS secrets to file")
	s_clientCmd.Flags().StringS("keymatexport", "keymatexport", "", "Export keying material using label")
	s_clientCmd.Flags().StringS("keymatexportlen", "keymatexportlen", "", "Export len bytes of keying material; default 20")
	s_clientCmd.Flags().BoolS("ktls", "ktls", false, "Enable Kernel TLS for sending and receiving")
	s_clientCmd.Flags().StringS("max_pipelines", "max_pipelines", "", "Maximum number of encrypt/decrypt pipelines to be used")
	s_clientCmd.Flags().StringS("max_send_frag", "max_send_frag", "", "Maximum Size of send frames")
	s_clientCmd.Flags().StringS("maxfraglen", "maxfraglen", "", "Enable Maximum Fragment Length Negotiation (len values: 512, 1024, 2048 and 4096)")
	s_clientCmd.Flags().BoolS("msg", "msg", false, "Show protocol messages")
	s_clientCmd.Flags().StringS("msgfile", "msgfile", "", "File to send output of -msg or -trace, instead of stdout")
	s_clientCmd.Flags().StringS("mtu", "mtu", "", "Set the link layer MTU")
	s_clientCmd.Flags().StringS("name", "name", "", "Hostname to use for \"-starttls lmtp\", \"-starttls smtp\" or \"-starttls xmpp[-server]\"")
	s_clientCmd.Flags().StringSliceS("nameopt", "nameopt", nil, "Certificate subject/issuer name printing options")
	s_clientCmd.Flags().BoolS("nbio", "nbio", false, "Use non-blocking IO")
	s_clientCmd.Flags().BoolS("nbio_test", "nbio_test", false, "More ssl protocol testing")
	s_clientCmd.Flags().StringS("nextprotoneg", "nextprotoneg", "", "Enable NPN extension, considering named protocols supported (comma-separated list)")
	s_clientCmd.Flags().BoolS("no-CAfile", "no-CAfile", false, "Do not load the default certificates file")
	s_clientCmd.Flags().BoolS("no-CApath", "no-CApath", false, "Do not load certificates from the default certificates directory")
	s_clientCmd.Flags().BoolS("no-CAstore", "no-CAstore", false, "Do not load certificates from the default certificates store")
	s_clientCmd.Flags().BoolS("no-interactive", "no-interactive", false, "Don't run the client in the interactive mode")
	s_clientCmd.Flags().BoolS("no_ign_eof", "no_ign_eof", false, "Don't ignore input eof")
	s_clientCmd.Flags().BoolS("nocommands", "nocommands", false, "Do not use interactive command letters")
	s_clientCmd.Flags().BoolS("noct", "noct", false, "Do not request or parse SCTs (default)")
	s_clientCmd.Flags().BoolS("noservername", "noservername", false, "Do not send the server name (SNI) extension in the ClientHello")
	s_clientCmd.Flags().BoolS("ocsp_check_all", "ocsp_check_all", false, "Require checking status of full chain, attempting to use OCSP stapling first")
	s_clientCmd.Flags().BoolS("ocsp_check_leaf", "ocsp_check_leaf", false, "Require checking leaf certificate status, attempting to use OCSP stapling first")
	s_clientCmd.Flags().StringS("pass", "pass", "", "Private key and cert file pass phrase source")
	s_clientCmd.Flags().StringS("port", "port", "", "Use -connect instead")
	s_clientCmd.Flags().BoolS("prexit", "prexit", false, "Print session information when the program exits")
	s_clientCmd.Flags().StringS("proxy", "proxy", "", "Connect to via specified proxy to the real server")
	s_clientCmd.Flags().StringS("proxy_pass", "proxy_pass", "", "Proxy authentication password source")
	s_clientCmd.Flags().StringS("proxy_user", "proxy_user", "", "UserID for proxy authentication")
	s_clientCmd.Flags().StringS("psk", "psk", "", "PSK in hex (without 0x)")
	s_clientCmd.Flags().StringS("psk_identity", "psk_identity", "", "PSK identity")
	s_clientCmd.Flags().StringS("psk_session", "psk_session", "", "File to read PSK SSL session from")
	s_clientCmd.Flags().BoolS("quic", "quic", false, "Use QUIC")
	s_clientCmd.Flags().BoolS("quiet", "quiet", false, "No s_client output")
	s_clientCmd.Flags().StringS("read_buf", "read_buf", "", "Default read buffer size to be used for connections")
	s_clientCmd.Flags().BoolS("reconnect", "reconnect", false, "Drop and re-make the connection with the same Session-ID")
	s_clientCmd.Flags().StringS("requestCAfile", "requestCAfile", "", "PEM format file of CA names to send to the server")
	s_clientCmd.Flags().BoolS("sctp", "sctp", false, "Use SCTP")
	s_clientCmd.Flags().BoolS("sctp_label_bug", "sctp_label_bug", false, "Enable SCTP label length bug")
	s_clientCmd.Flags().BoolS("security_debug", "security_debug", false, "Enable security debug messages")
	s_clientCmd.Flags().BoolS("security_debug_verbose", "security_debug_verbose", false, "Output more security debug output")
	s_clientCmd.Flags().StringS("serverinfo", "serverinfo", "", "types  Send empty ClientHello extensions (comma-separated numbers)")
	s_clientCmd.Flags().StringS("servername", "servername", "", "Set TLS extension servername (SNI) in ClientHello (default)")
	s_clientCmd.Flags().StringS("sess_in", "sess_in", "", "File to read SSL session from")
	s_clientCmd.Flags().StringS("sess_out", "sess_out", "", "File to write SSL session to")
	s_clientCmd.Flags().BoolS("showcerts", "showcerts", false, "Show all certificates sent by the server")
	s_clientCmd.Flags().StringS("split_send_frag", "split_send_frag", "", "Size used to split data for encrypt pipelines")
	s_clientCmd.Flags().BoolS("srp_lateuser", "srp_lateuser", false, "(deprecated) SRP username into second ClientHello message")
	s_clientCmd.Flags().BoolS("srp_moregroups", "srp_moregroups", false, "(deprecated) Tolerate other than the known g N values.")
	s_clientCmd.Flags().StringS("srp_strength", "srp_strength", "", "(deprecated) Minimal length in bits for N")
	s_clientCmd.Flags().StringS("srppass", "srppass", "", "(deprecated) Password for 'user'")
	s_clientCmd.Flags().StringS("srpuser", "srpuser", "", "(deprecated) SRP authentication for 'user'")
	s_clientCmd.Flags().StringS("ssl_config", "ssl_config", "", "Use specified section for SSL_CTX configuration")
	s_clientCmd.Flags().StringS("starttls", "starttls", "", "Use the appropriate STARTTLS command before starting TLS")
	s_clientCmd.Flags().BoolS("state", "state", false, "Print the ssl states")
	s_clientCmd.Flags().BoolS("status", "status", false, "Sends a certificate status request to the server (OCSP stapling) The server response (if any) will be printed out.")
	s_clientCmd.Flags().BoolS("tfo", "tfo", false, "Connect using TCP Fast Open")
	s_clientCmd.Flags().BoolS("timeout", "timeout", false, "Enable send/receive timeout on DTLS connections")
	s_clientCmd.Flags().BoolS("tls1", "tls1", false, "Just use TLSv1")
	s_clientCmd.Flags().BoolS("tls1_1", "tls1_1", false, "Just use TLSv1.1")
	s_clientCmd.Flags().BoolS("tls1_2", "tls1_2", false, "Just use TLSv1.2")
	s_clientCmd.Flags().BoolS("tls1_3", "tls1_3", false, "Just use TLSv1.3")
	s_clientCmd.Flags().BoolS("tlsextdebug", "tlsextdebug", false, "Hex dump of all TLS extensions received")
	s_clientCmd.Flags().BoolS("trace", "trace", false, "Show trace output of protocol messages")
	s_clientCmd.Flags().StringS("unix", "unix", "", "Connect over the specified Unix-domain socket")
	s_clientCmd.Flags().StringS("use_srtp", "use_srtp", "", "Offer SRTP key management with a colon-separated profile list")
	s_clientCmd.Flags().StringS("verify", "verify", "", "Turn on peer certificate verification")
	s_clientCmd.Flags().StringS("verifyCAfile", "verifyCAfile", "", "CA file for certificate verification (PEM format)")
	s_clientCmd.Flags().StringS("verifyCApath", "verifyCApath", "", "Use dir as certificate store path to verify CA certificate")
	s_clientCmd.Flags().StringS("verifyCAstore", "verifyCAstore", "", "CA store URI for certificate verification")
	s_clientCmd.Flags().BoolS("verify_quiet", "verify_quiet", false, "Restrict verify output to errors")
	s_clientCmd.Flags().BoolS("verify_return_error", "verify_return_error", false, "Close connection on verification error")
	s_clientCmd.Flags().BoolS("wdebug", "wdebug", false, "WATT-32 tcp debugging")
	s_clientCmd.Flags().StringS("xmpphost", "xmpphost", "", "Alias of -name option for \"-starttls xmpp[-server]\"")
	common.AddExtendedCertificateOptions(s_clientCmd)
	common.AddProviderFlags(s_clientCmd)
	common.AddRandomStateFlags(s_clientCmd)
	common.AddTLSAndSSLOptions(s_clientCmd)
	common.AddValidationFlags(s_clientCmd)
	rootCmd.AddCommand(s_clientCmd)

	carapace.Gen(s_clientCmd).FlagCompletion(carapace.ActionMap{
		"CAfile":        carapace.ActionFiles(),
		"CApath":        carapace.ActionDirectories(),
		"CRL":           carapace.ActionFiles(),
		"CRLform":       carapace.ActionValues("DER", "PEM"),
		"cert":          carapace.ActionFiles(),
		"certform":      carapace.ActionValues("DER", "PEM", "P12"),
		"chainCAfile":   carapace.ActionFiles(),
		"chainCApath":   carapace.ActionDirectories(),
		"early_data":    carapace.ActionFiles(),
		"key":           carapace.ActionFiles(),
		"keyform":       carapace.ActionValues("DER", "PEM", "P12"),
		"keylogfile":    carapace.ActionFiles(),
		"msgfile":       carapace.ActionFiles(),
		"psk_session":   carapace.ActionFiles(),
		"requestCAfile": carapace.ActionFiles(),
		"sess_in":       carapace.ActionFiles(),
		"sess_out":      carapace.ActionFiles(),
		"verifyCAfile":  carapace.ActionFiles(),
		"verifyCApath":  carapace.ActionDirectories(),
	})
}
