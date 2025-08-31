package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var s_serverCmd = &cobra.Command{
	Use:     "s_server",
	Short:   "SSL/TLS server program",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s_serverCmd).Standalone()

	s_serverCmd.Flags().BoolS("4", "4", false, "Use IPv4 only")
	s_serverCmd.Flags().BoolS("6", "6", false, "Use IPv6 only")
	s_serverCmd.Flags().StringS("CAfile", "CAfile", "", "PEM format file of CA's")
	s_serverCmd.Flags().StringS("CApath", "CApath", "", "PEM format directory of CA's")
	s_serverCmd.Flags().StringS("CAstore", "CAstore", "", "URI to store of CA's")
	s_serverCmd.Flags().StringS("CRL", "CRL", "", "CRL file to use")
	s_serverCmd.Flags().StringS("CRLform", "CRLform", "", "CRL file format (PEM or DER); default PEM")
	s_serverCmd.Flags().BoolS("HTTP", "HTTP", false, "Like -WWW but ./path includes HTTP headers")
	s_serverCmd.Flags().StringS("Verify", "Verify", "", "Turn on peer certificate verification, must have a cert")
	s_serverCmd.Flags().BoolS("WWW", "WWW", false, "Respond to a 'GET with the file ./path")
	s_serverCmd.Flags().StringS("accept", "accept", "", "TCP/IP optional host and port to listen on for connections (default is *:4433)")
	s_serverCmd.Flags().StringS("alpn", "alpn", "", "Set the advertised protocols for the ALPN extension (comma-separated list)")
	s_serverCmd.Flags().BoolS("anti_replay", "anti_replay", false, "Switch on anti-replay protection (default)")
	s_serverCmd.Flags().BoolS("async", "async", false, "Operate in asynchronous mode")
	s_serverCmd.Flags().BoolS("brief", "brief", false, "Restrict output to brief summary of connection parameters")
	s_serverCmd.Flags().BoolS("build_chain", "build_chain", false, "Build server certificate chain")
	s_serverCmd.Flags().StringS("cert", "cert", "", "Server certificate file to use; default server.pem")
	s_serverCmd.Flags().StringS("cert2", "cert2", "", "Certificate file to use for servername; default server2.pem")
	s_serverCmd.Flags().StringS("cert_chain", "cert_chain", "", "Server certificate chain file in PEM format")
	s_serverCmd.Flags().BoolS("cert_comp", "cert_comp", false, "Pre-compress server certificates")
	s_serverCmd.Flags().StringS("certform", "certform", "", "Server certificate file format (PEM/DER/P12); has no effect")
	s_serverCmd.Flags().StringS("chainCAfile", "chainCAfile", "", "CA file for certificate chain (PEM format)")
	s_serverCmd.Flags().StringS("chainCApath", "chainCApath", "", "use dir as certificate store path to build CA certificate chain")
	s_serverCmd.Flags().StringS("chainCAstore", "chainCAstore", "", "use URI as certificate store to build CA certificate chain")
	s_serverCmd.Flags().StringS("context", "context", "", "Set session ID context")
	s_serverCmd.Flags().BoolS("crl_download", "crl_download", false, "Download CRLs from distribution points in certificate CDP entries")
	s_serverCmd.Flags().BoolS("crlf", "crlf", false, "Convert LF from terminal into CRLF")
	s_serverCmd.Flags().StringS("dcert", "dcert", "", "Second server certificate file to use (usually for DSA)")
	s_serverCmd.Flags().StringS("dcert_chain", "dcert_chain", "", "second server certificate chain file in PEM format")
	s_serverCmd.Flags().StringS("dcertform", "dcertform", "", "Second server certificate file format (PEM/DER/P12); has no effect")
	s_serverCmd.Flags().BoolS("debug", "debug", false, "Print more output")
	s_serverCmd.Flags().StringS("dhparam", "dhparam", "", "DH parameters file to use")
	s_serverCmd.Flags().StringS("dkey", "dkey", "", "Second private key file to use (usually for DSA)")
	s_serverCmd.Flags().StringS("dkeyform", "dkeyform", "", "Second key file format (ENGINE, other values ignored)")
	s_serverCmd.Flags().StringS("dpass", "dpass", "", "Second private key and cert file pass phrase source")
	s_serverCmd.Flags().BoolS("dtls", "dtls", false, "Use any DTLS version")
	s_serverCmd.Flags().BoolS("dtls1", "dtls1", false, "Just talk DTLSv1")
	s_serverCmd.Flags().BoolS("dtls1_2", "dtls1_2", false, "Just talk DTLSv1.2")
	s_serverCmd.Flags().BoolS("early_data", "early_data", false, "Attempt to read early data")
	s_serverCmd.Flags().BoolS("enable_client_rpk", "enable_client_rpk", false, "Enable raw public keys (RFC7250) from the client")
	s_serverCmd.Flags().BoolS("enable_server_rpk", "enable_server_rpk", false, "Enable raw public keys (RFC7250) from the server")
	s_serverCmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	s_serverCmd.Flags().BoolS("ext_cache", "ext_cache", false, "Disable internal cache, set up and use external cache")
	s_serverCmd.Flags().BoolS("http_server_binmode", "http_server_binmode", false, "opening files in binary mode when acting as http server (-WWW and -HTTP)")
	s_serverCmd.Flags().StringS("id_prefix", "id_prefix", "", "Generate SSL/TLS session IDs prefixed by arg")
	s_serverCmd.Flags().BoolS("ign_eof", "ign_eof", false, "Ignore input EOF (default when -quiet)")
	s_serverCmd.Flags().BoolS("ignore_unexpected_eof", "ignore_unexpected_eof", false, "Do not treat lack of close_notify from a peer as an error")
	s_serverCmd.Flags().StringS("key", "key", "", "Private key file to use; default is -cert file or elseserver.pem")
	s_serverCmd.Flags().StringS("key2", "key2", "", "-Private Key file to use for servername if not in -cert2")
	s_serverCmd.Flags().StringS("keyform", "keyform", "", "Key format (ENGINE, other values ignored)")
	s_serverCmd.Flags().StringS("keylogfile", "keylogfile", "", "Write TLS secrets to file")
	s_serverCmd.Flags().StringS("keymatexport", "keymatexport", "", "Export keying material using label")
	s_serverCmd.Flags().StringS("keymatexportlen", "keymatexportlen", "", "Export len bytes of keying material; default 20")
	s_serverCmd.Flags().BoolS("ktls", "ktls", false, "Enable Kernel TLS for sending and receiving")
	s_serverCmd.Flags().BoolS("listen", "listen", false, "Listen for a DTLS ClientHello with a cookie and then connect")
	s_serverCmd.Flags().StringS("max_early_data", "max_early_data", "", "The maximum number of bytes of early data as advertised in tickets")
	s_serverCmd.Flags().StringS("max_pipelines", "max_pipelines", "", "Maximum number of encrypt/decrypt pipelines to be used")
	s_serverCmd.Flags().StringS("max_send_frag", "max_send_frag", "", "Maximum Size of send frames")
	s_serverCmd.Flags().BoolS("msg", "msg", false, "Show protocol messages")
	s_serverCmd.Flags().StringS("msgfile", "msgfile", "", "File to send output of -msg or -trace, instead of stdout")
	s_serverCmd.Flags().StringS("mtu", "mtu", "", "Set link-layer MTU")
	s_serverCmd.Flags().StringS("naccept", "naccept", "", "Terminate after #num connections")
	s_serverCmd.Flags().StringSliceS("nameopt", "nameopt", nil, "Certificate subject/issuer name printing options")
	s_serverCmd.Flags().BoolS("nbio", "nbio", false, "Use non-blocking IO")
	s_serverCmd.Flags().BoolS("nbio_test", "nbio_test", false, "Test with the non-blocking test bio")
	s_serverCmd.Flags().StringS("nextprotoneg", "nextprotoneg", "", "Set the advertised protocols for the NPN extension (comma-separated list)")
	s_serverCmd.Flags().BoolS("no-CAfile", "no-CAfile", false, "Do not load the default certificates file")
	s_serverCmd.Flags().BoolS("no-CApath", "no-CApath", false, "Do not load certificates from the default certificates directory")
	s_serverCmd.Flags().BoolS("no-CAstore", "no-CAstore", false, "Do not load certificates from the default certificates store URI")
	s_serverCmd.Flags().BoolS("no_anti_replay", "no_anti_replay", false, "Switch off anti-replay protection")
	s_serverCmd.Flags().BoolS("no_ca_names", "no_ca_names", false, "Disable TLS Extension CA Names")
	s_serverCmd.Flags().BoolS("no_cache", "no_cache", false, "Disable session cache")
	s_serverCmd.Flags().BoolS("no_dhe", "no_dhe", false, "Disable ephemeral DH")
	s_serverCmd.Flags().BoolS("no_ign_eof", "no_ign_eof", false, "Do not ignore input EOF")
	s_serverCmd.Flags().StringS("no_proxy", "no_proxy", "", "List of addresses of servers not to use HTTP(S) proxy for Default from environment variable 'no_proxy', else 'NO_PROXY', else none")
	s_serverCmd.Flags().BoolS("no_resume_ephemeral", "no_resume_ephemeral", false, "Disable caching and tickets if ephemeral (EC)DH is used")
	s_serverCmd.Flags().BoolS("nocert", "nocert", false, "Don't use any certificates (Anon-DH)")
	s_serverCmd.Flags().StringS("num_tickets", "num_tickets", "", "The number of TLSv1.3 session tickets that a server will automatically issue")
	s_serverCmd.Flags().StringS("pass", "pass", "", "Private key and cert file pass phrase source")
	s_serverCmd.Flags().StringS("port", "port", "", "TCP/IP port to listen on for connections (default is 4433)")
	s_serverCmd.Flags().StringS("proxy", "proxy", "", "[http[s]://]host[:port][/path] of HTTP(S) proxy to use; path is ignored")
	s_serverCmd.Flags().StringS("psk", "psk", "", "PSK in hex (without 0x)")
	s_serverCmd.Flags().StringS("psk_hint", "psk_hint", "", "PSK identity hint to use")
	s_serverCmd.Flags().StringS("psk_identity", "psk_identity", "", "PSK identity to expect")
	s_serverCmd.Flags().StringS("psk_session", "psk_session", "", "File to read PSK SSL session from")
	s_serverCmd.Flags().BoolS("quiet", "quiet", false, "No server output")
	s_serverCmd.Flags().StringS("read_buf", "read_buf", "", "Default read buffer size to be used for connections")
	s_serverCmd.Flags().StringS("recv_max_early_data", "recv_max_early_data", "", "The maximum number of bytes of early data (hard limit)")
	s_serverCmd.Flags().BoolS("rev", "rev", false, "act as an echo server that sends back received text reversed")
	s_serverCmd.Flags().BoolS("sctp", "sctp", false, "Use SCTP")
	s_serverCmd.Flags().BoolS("sctp_label_bug", "sctp_label_bug", false, "Enable SCTP label length bug")
	s_serverCmd.Flags().BoolS("security_debug", "security_debug", false, "Print output from SSL/TLS security framework")
	s_serverCmd.Flags().BoolS("security_debug_verbose", "security_debug_verbose", false, "Print more output from SSL/TLS security framework")
	s_serverCmd.Flags().BoolS("sendfile", "sendfile", false, "Use sendfile to response file with -WWW")
	s_serverCmd.Flags().StringS("serverinfo", "serverinfo", "", "PEM serverinfo file for certificate")
	s_serverCmd.Flags().StringS("servername", "servername", "", "Servername for HostName TLS extension")
	s_serverCmd.Flags().BoolS("servername_fatal", "servername_fatal", false, "On servername mismatch send fatal alert (default warning alert)")
	s_serverCmd.Flags().StringS("split_send_frag", "split_send_frag", "", "Size used to split data for encrypt pipelines")
	s_serverCmd.Flags().StringS("srpuserseed", "srpuserseed", "", "(deprecated) A seed string for a default user salt")
	s_serverCmd.Flags().StringS("srpvfile", "srpvfile", "", "(deprecated) The verifier file for SRP")
	s_serverCmd.Flags().BoolS("ssl3", "ssl3", false, "Just talk SSLv3")
	s_serverCmd.Flags().StringS("ssl_config", "ssl_config", "", "Configure SSL_CTX using the given configuration value")
	s_serverCmd.Flags().BoolS("state", "state", false, "Print the SSL states")
	s_serverCmd.Flags().BoolS("stateless", "stateless", false, "Require TLSv1.3 cookies")
	s_serverCmd.Flags().BoolS("status", "status", false, "Provide certificate status response if requested, for server cert only")
	s_serverCmd.Flags().BoolS("status_all", "status_all", false, "Provide certificate status response(s) if requested, for the whole chain")
	s_serverCmd.Flags().StringSliceS("status_file", "status_file", nil, "File containing DER encoded OCSP Response (can be specified multiple times)")
	s_serverCmd.Flags().StringS("status_timeout", "status_timeout", "", "Status request responder timeout")
	s_serverCmd.Flags().StringS("status_url", "status_url", "", "Status request fallback URL")
	s_serverCmd.Flags().BoolS("status_verbose", "status_verbose", false, "Print more output in certificate status callback")
	s_serverCmd.Flags().BoolS("tfo", "tfo", false, "Listen for TCP Fast Open connections")
	s_serverCmd.Flags().BoolS("timeout", "timeout", false, "Enable timeouts")
	s_serverCmd.Flags().BoolS("tls1", "tls1", false, "Just talk TLSv1")
	s_serverCmd.Flags().BoolS("tls1_1", "tls1_1", false, "Just talk TLSv1.1")
	s_serverCmd.Flags().BoolS("tls1_2", "tls1_2", false, "just talk TLSv1.2")
	s_serverCmd.Flags().BoolS("tls1_3", "tls1_3", false, "just talk TLSv1.3")
	s_serverCmd.Flags().BoolS("tlsextdebug", "tlsextdebug", false, "Hex dump of all TLS extensions received")
	s_serverCmd.Flags().BoolS("trace", "trace", false, "trace protocol messages")
	s_serverCmd.Flags().StringS("unix", "unix", "", "Unix domain socket to accept on")
	s_serverCmd.Flags().BoolS("unlink", "unlink", false, "For -unix, unlink existing socket first")
	s_serverCmd.Flags().StringS("use_srtp", "use_srtp", "", "Offer SRTP key management with a colon-separated profile list")
	s_serverCmd.Flags().StringS("verify", "verify", "", "Turn on peer certificate verification")
	s_serverCmd.Flags().StringS("verifyCAfile", "verifyCAfile", "", "CA file for certificate verification (PEM format)")
	s_serverCmd.Flags().StringS("verifyCApath", "verifyCApath", "", "use dir as certificate store path to verify CA certificate")
	s_serverCmd.Flags().StringS("verifyCAstore", "verifyCAstore", "", "use URI as certificate store to verify CA certificate")
	s_serverCmd.Flags().BoolS("verify_quiet", "verify_quiet", false, "No verify output except verify errors")
	s_serverCmd.Flags().BoolS("verify_return_error", "verify_return_error", false, "Close connection on verification error")
	s_serverCmd.Flags().BoolS("www", "www", false, "Respond to a 'GET /' with a status page")
	s_serverCmd.Flags().BoolS("zerocopy_sendfile", "zerocopy_sendfile", false, "Use zerocopy mode of KTLS sendfile")
	common.AddExtendedCertificateOptions(s_serverCmd)
	common.AddProviderFlags(s_serverCmd)
	common.AddRandomStateFlags(s_serverCmd)
	common.AddTLSAndSSLOptions(s_serverCmd)
	common.AddValidationFlags(s_serverCmd)
	rootCmd.AddCommand(s_serverCmd)

	s_serverCmd.MarkFlagsMutuallyExclusive("tls1_3", "nextprotoneg")

	carapace.Gen(s_serverCmd).FlagCompletion(carapace.ActionMap{
		"CAfile":       carapace.ActionFiles(),
		"CApath":       carapace.ActionDirectories(),
		"CRL":          carapace.ActionFiles(),
		"CRLform":      carapace.ActionValues("DER", "PEM"),
		"cert":         carapace.ActionFiles(),
		"cert2":        carapace.ActionFiles(),
		"cert_chain":   carapace.ActionFiles(),
		"certform":     carapace.ActionValues("DER", "PEM", "P12"),
		"chainCAfile":  carapace.ActionFiles(),
		"chainCApath":  carapace.ActionDirectories(),
		"dcert":        carapace.ActionFiles(),
		"dcert_chain":  carapace.ActionFiles(),
		"dcertform":    carapace.ActionValues("DER", "PEM", "P12"),
		"dkey":         carapace.ActionFiles(),
		"dkeyform":     carapace.ActionValues("ENGINE", "DER", "PEM", "P12"),
		"engine":       action.ActionEngines(),
		"key":          carapace.ActionFiles(),
		"key2":         carapace.ActionFiles(),
		"keyform":      carapace.ActionValues("ENGINE", "DER", "PEM", "P12"),
		"keylogfile":   carapace.ActionFiles(),
		"msgfile":      carapace.ActionFiles(),
		"psk_session":  carapace.ActionFiles(),
		"srpvfile":     carapace.ActionFiles(),
		"status_file":  carapace.ActionFiles(),
		"verifyCAfile": carapace.ActionFiles(),
		"verifyCApath": carapace.ActionDirectories(),
	})
}
