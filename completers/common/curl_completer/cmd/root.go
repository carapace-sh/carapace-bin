package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/curl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "curl",
	Short: "transfer a URL",
	Long:  "https://github.com/curl/curl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("abstract-unix-socket", "", "Connect via abstract Unix domain socket")
	rootCmd.Flags().String("alt-svc", "", "Enable alt-svc with this cache file")
	rootCmd.Flags().Bool("anyauth", false, "Pick any authentication method")
	rootCmd.Flags().BoolP("append", "a", false, "Append to target file when uploading")
	rootCmd.Flags().String("aws-sigv4", "", "AWS V4 signature auth")
	rootCmd.Flags().Bool("basic", false, "HTTP Basic Authentication")
	rootCmd.Flags().Bool("ca-native", false, "Load CA certs from the OS")
	rootCmd.Flags().String("cacert", "", "CA certificate to verify peer against")
	rootCmd.Flags().String("capath", "", "CA directory to verify peer against")
	rootCmd.Flags().StringP("cert", "E", "", "Client certificate file and password")
	rootCmd.Flags().Bool("cert-status", false, "Verify server cert status OCSP-staple")
	rootCmd.Flags().String("cert-type", "", "Certificate type (DER/PEM/ENG/PROV/P12)")
	rootCmd.Flags().String("ciphers", "", "TLS 1.2 (1.1, 1.0) ciphers to use")
	rootCmd.Flags().Bool("compressed", false, "Request compressed response")
	rootCmd.Flags().String("compressed-ssh", "", "Enable SSH compression")
	rootCmd.Flags().StringP("config", "K", "", "Read config from a file")
	rootCmd.Flags().String("connect-timeout", "", "Maximum time allowed to connect")
	rootCmd.Flags().String("connect-to", "", "Connect to host2 instead of host1")
	rootCmd.Flags().StringP("continue-at", "C", "", "Resumed transfer offset")
	rootCmd.Flags().StringP("cookie", "b", "", "Send cookies from string/load from file")
	rootCmd.Flags().StringP("cookie-jar", "c", "", "Save cookies to <filename> after operation")
	rootCmd.Flags().Bool("create-dirs", false, "Create necessary local directory hierarchy")
	rootCmd.Flags().String("create-file-mode", "", "File mode for created files")
	rootCmd.Flags().Bool("crlf", false, "Convert LF to CRLF in upload")
	rootCmd.Flags().String("crlfile", "", "Certificate Revocation list")
	rootCmd.Flags().String("curves", "", "(EC) TLS key exchange algorithms to request")
	rootCmd.Flags().StringP("data", "d", "", "HTTP POST data")
	rootCmd.Flags().String("data-ascii", "", "HTTP POST ASCII data")
	rootCmd.Flags().String("data-binary", "", "HTTP POST binary data")
	rootCmd.Flags().String("data-raw", "", "HTTP POST data, '@' allowed")
	rootCmd.Flags().String("data-urlencode", "", "HTTP POST data URL encoded")
	rootCmd.Flags().String("delegation", "", "GSS-API delegation permission")
	rootCmd.Flags().Bool("digest", false, "HTTP Digest Authentication")
	rootCmd.Flags().BoolP("disable", "q", false, "Disable .curlrc")
	rootCmd.Flags().Bool("disable-eprt", false, "Inhibit using EPRT or LPRT")
	rootCmd.Flags().Bool("disable-epsv", false, "Inhibit using EPSV")
	rootCmd.Flags().String("disallow-username-in-url", "", "Disallow username in URL")
	rootCmd.Flags().String("dns-interface", "", "Interface to use for DNS requests")
	rootCmd.Flags().String("dns-ipv4-addr", "", "IPv4 address to use for DNS requests")
	rootCmd.Flags().String("dns-ipv6-addr", "", "IPv6 address to use for DNS requests")
	rootCmd.Flags().String("dns-servers", "", "DNS server addrs to use")
	rootCmd.Flags().Bool("doh-cert-status", false, "Verify DoH server cert status OCSP-staple")
	rootCmd.Flags().Bool("doh-insecure", false, "Allow insecure DoH server connections")
	rootCmd.Flags().String("doh-url", "", "Resolve hostnames over DoH")
	rootCmd.Flags().Bool("dump-ca-embed", false, "Write the embedded CA bundle to standard output")
	rootCmd.Flags().StringP("dump-header", "D", "", "Write the received headers to <filename>")
	rootCmd.Flags().String("ech", "", "Configure ECH")
	rootCmd.Flags().String("egd-file", "", "EGD socket path for random data")
	rootCmd.Flags().String("engine", "", "Crypto engine to use")
	rootCmd.Flags().String("etag-compare", "", "Load ETag from file")
	rootCmd.Flags().String("etag-save", "", "Parse incoming ETag and save to a file")
	rootCmd.Flags().String("expect100-timeout", "", "How long to wait for 100-continue")
	rootCmd.Flags().BoolP("fail", "f", false, "Fail fast with no output on HTTP errors")
	rootCmd.Flags().Bool("fail-early", false, "Fail on first transfer error")
	rootCmd.Flags().Bool("fail-with-body", false, "Fail on HTTP errors but save the body")
	rootCmd.Flags().Bool("false-start", false, "Enable TLS False Start")
	rootCmd.Flags().Bool("follow", false, "Follow redirects per spec")
	rootCmd.Flags().StringP("form", "F", "", "Specify multipart MIME data")
	rootCmd.Flags().Bool("form-escape", false, "Escape form fields using backslash")
	rootCmd.Flags().String("form-string", "", "Specify multipart MIME data")
	rootCmd.Flags().String("ftp-account", "", "Account data string")
	rootCmd.Flags().String("ftp-alternative-to-user", "", "String to replace USER [name]")
	rootCmd.Flags().String("ftp-create-dirs", "", "Create the remote dirs if not present")
	rootCmd.Flags().String("ftp-method", "", "Control CWD usage")
	rootCmd.Flags().Bool("ftp-pasv", false, "Send PASV/EPSV instead of PORT")
	rootCmd.Flags().StringP("ftp-port", "P", "", "Send PORT instead of PASV")
	rootCmd.Flags().Bool("ftp-pret", false, "Send PRET before PASV")
	rootCmd.Flags().String("ftp-skip-pasv-ip", "", "Skip the IP address for PASV")
	rootCmd.Flags().Bool("ftp-ssl-ccc", false, "Send CCC after authenticating")
	rootCmd.Flags().String("ftp-ssl-ccc-mode", "", "Set CCC mode")
	rootCmd.Flags().String("ftp-ssl-control", "", "Require TLS for login, clear for transfer")
	rootCmd.Flags().BoolP("get", "G", false, "Put the post data in the URL and use GET")
	rootCmd.Flags().BoolP("globoff", "g", false, "Disable URL globbing with {} and []")
	rootCmd.Flags().String("happy-eyeballs-timeout-ms", "", "Time for IPv6 before IPv4")
	rootCmd.Flags().String("haproxy-clientip", "", "Set address in HAProxy PROXY")
	rootCmd.Flags().String("haproxy-protocol", "", "Send HAProxy PROXY protocol v1 header")
	rootCmd.Flags().BoolP("head", "I", false, "Show document info only")
	rootCmd.Flags().StringArrayP("header", "H", nil, "Pass custom header(s) to server")
	rootCmd.Flags().BoolP("help", "h", false, "Get help for commands")
	rootCmd.Flags().String("hostpubmd5", "", "Acceptable MD5 hash of host public key")
	rootCmd.Flags().String("hostpubsha256", "", "Acceptable SHA256 hash of host public key")
	rootCmd.Flags().String("hsts", "", "Enable HSTS with this cache file")
	rootCmd.Flags().Bool("http0.9", false, "Allow HTTP/0.9 responses")
	rootCmd.Flags().BoolP("http1.0", "0", false, "Use HTTP/1.0")
	rootCmd.Flags().Bool("http1.1", false, "Use HTTP/1.1")
	rootCmd.Flags().Bool("http2", false, "Use HTTP/2")
	rootCmd.Flags().String("http2-prior-knowledge", "", "Use HTTP/2 without HTTP/1.1 Upgrade")
	rootCmd.Flags().Bool("http3", false, "Use HTTP/3")
	rootCmd.Flags().Bool("http3-only", false, "Use HTTP/3 only")
	rootCmd.Flags().String("ignore-content-length", "", "Ignore the size of the remote resource")
	rootCmd.Flags().BoolP("insecure", "k", false, "Allow insecure server connections")
	rootCmd.Flags().String("interface", "", "Use network interface")
	rootCmd.Flags().String("ip-tos", "", "Set IP Type of Service or Traffic Class")
	rootCmd.Flags().String("ipfs-gateway", "", "Gateway for IPFS")
	rootCmd.Flags().BoolP("ipv4", "4", false, "Resolve names to IPv4 addresses")
	rootCmd.Flags().BoolP("ipv6", "6", false, "Resolve names to IPv6 addresses")
	rootCmd.Flags().String("json", "", "HTTP POST JSON")
	rootCmd.Flags().StringP("junk-session-cookies", "j", "", "Ignore session cookies read from file")
	rootCmd.Flags().String("keepalive-cnt", "", "Maximum number of keepalive probes")
	rootCmd.Flags().String("keepalive-time", "", "Interval time for keepalive probes")
	rootCmd.Flags().String("key", "", "Private key filename")
	rootCmd.Flags().String("key-type", "", "Private key file type (DER/PEM/ENG)")
	rootCmd.Flags().String("krb", "", "Enable Kerberos with security <level>")
	rootCmd.Flags().String("libcurl", "", "Generate libcurl code for this command line")
	rootCmd.Flags().String("limit-rate", "", "Limit transfer speed to RATE")
	rootCmd.Flags().BoolP("list-only", "l", false, "List only mode")
	rootCmd.Flags().String("local-port", "", "Use a local port number within RANGE")
	rootCmd.Flags().BoolP("location", "L", false, "Follow redirects")
	rootCmd.Flags().String("location-trusted", "", "As --location, but send secrets to other hosts")
	rootCmd.Flags().String("login-options", "", "Server login options")
	rootCmd.Flags().String("mail-auth", "", "Originator address of the original email")
	rootCmd.Flags().String("mail-from", "", "Mail from this address")
	rootCmd.Flags().String("mail-rcpt", "", "Mail to this address")
	rootCmd.Flags().String("mail-rcpt-allowfails", "", "Allow RCPT TO command to fail")
	rootCmd.Flags().BoolP("manual", "M", false, "Display the full manual")
	rootCmd.Flags().String("max-filesize", "", "Maximum file size to download")
	rootCmd.Flags().String("max-redirs", "", "Maximum number of redirects allowed")
	rootCmd.Flags().StringP("max-time", "m", "", "Maximum time allowed for transfer")
	rootCmd.Flags().Bool("metalink", false, "Process given URLs as metalink XML file")
	rootCmd.Flags().Bool("mptcp", false, "Enable Multipath TCP")
	rootCmd.Flags().Bool("negotiate", false, "Use HTTP Negotiate (SPNEGO) authentication")
	rootCmd.Flags().BoolP("netrc", "n", false, "Must read .netrc for username and password")
	rootCmd.Flags().String("netrc-file", "", "Specify FILE for netrc")
	rootCmd.Flags().String("netrc-optional", "", "Use either .netrc or URL")
	rootCmd.Flags().BoolP("next", ":", false, "Make next URL use separate options")
	rootCmd.Flags().Bool("no-alpn", false, "Disable the ALPN TLS extension")
	rootCmd.Flags().BoolP("no-buffer", "N", false, "Disable buffering of the output stream")
	rootCmd.Flags().Bool("no-clobber", false, "Do not overwrite files that already exist")
	rootCmd.Flags().Bool("no-keepalive", false, "Disable TCP keepalive on the connection")
	rootCmd.Flags().Bool("no-npn", false, "Disable the NPN TLS extension")
	rootCmd.Flags().String("no-progress-meter", "", "Do not show the progress meter")
	rootCmd.Flags().Bool("no-sessionid", false, "Disable SSL session-ID reusing")
	rootCmd.Flags().String("noproxy", "", "List of hosts which do not use proxy")
	rootCmd.Flags().Bool("ntlm", false, "HTTP NTLM authentication")
	rootCmd.Flags().Bool("ntlm-wb", false, "HTTP NTLM authentication with winbind")
	rootCmd.Flags().String("oauth2-bearer", "", "OAuth 2 Bearer Token")
	rootCmd.Flags().Bool("out-null", false, "Discard response data into the void")
	rootCmd.Flags().StringP("output", "o", "", "Write to file instead of stdout")
	rootCmd.Flags().String("output-dir", "", "Directory to save files in")
	rootCmd.Flags().BoolP("parallel", "Z", false, "Perform transfers in parallel")
	rootCmd.Flags().String("parallel-immediate", "", "Do not wait for multiplexing")
	rootCmd.Flags().Bool("parallel-max", false, "Maximum concurrency for parallel transfers")
	rootCmd.Flags().Bool("parallel-max-host", false, "Maximum connections to a single host")
	rootCmd.Flags().String("pass", "", "Passphrase for the private key")
	rootCmd.Flags().Bool("path-as-is", false, "Do not squash .. sequences in URL path")
	rootCmd.Flags().String("pinnedpubkey", "", "Public key to verify peer against")
	rootCmd.Flags().Bool("post301", false, "Do not switch to GET after a 301 redirect")
	rootCmd.Flags().Bool("post302", false, "Do not switch to GET after a 302 redirect")
	rootCmd.Flags().Bool("post303", false, "Do not switch to GET after a 303 redirect")
	rootCmd.Flags().String("preproxy", "", "Use this proxy first")
	rootCmd.Flags().BoolP("progress-bar", "#", false, "Display transfer progress as a bar")
	rootCmd.Flags().String("proto", "", "Enable/disable PROTOCOLS")
	rootCmd.Flags().String("proto-default", "", "Use PROTOCOL for any URL missing a scheme")
	rootCmd.Flags().String("proto-redir", "", "Enable/disable PROTOCOLS on redirect")
	rootCmd.Flags().StringP("proxy", "x", "", "Use this proxy")
	rootCmd.Flags().String("proxy-anyauth", "", "Pick any proxy authentication method")
	rootCmd.Flags().Bool("proxy-basic", false, "Use Basic authentication on the proxy")
	rootCmd.Flags().Bool("proxy-ca-native", false, "Load CA certs from the OS to verify proxy")
	rootCmd.Flags().String("proxy-cacert", "", "CA certificates to verify proxy against")
	rootCmd.Flags().String("proxy-capath", "", "CA directory to verify proxy against")
	rootCmd.Flags().String("proxy-cert", "", "Set client certificate for proxy")
	rootCmd.Flags().String("proxy-cert-type", "", "Client certificate type for HTTPS proxy")
	rootCmd.Flags().String("proxy-ciphers", "", "TLS 1.2 (1.1, 1.0) ciphers to use for proxy")
	rootCmd.Flags().String("proxy-crlfile", "", "Set a CRL list for proxy")
	rootCmd.Flags().Bool("proxy-digest", false, "Digest auth with the proxy")
	rootCmd.Flags().String("proxy-header", "", "Pass custom header(s) to proxy")
	rootCmd.Flags().Bool("proxy-http2", false, "Use HTTP/2 with HTTPS proxy")
	rootCmd.Flags().String("proxy-insecure", "", "Skip HTTPS proxy cert verification")
	rootCmd.Flags().String("proxy-key", "", "Private key for HTTPS proxy")
	rootCmd.Flags().String("proxy-key-type", "", "Private key file type for proxy")
	rootCmd.Flags().String("proxy-negotiate", "", "HTTP Negotiate (SPNEGO) auth with the proxy")
	rootCmd.Flags().Bool("proxy-ntlm", false, "NTLM authentication with the proxy")
	rootCmd.Flags().String("proxy-pass", "", "Passphrase for private key for HTTPS proxy")
	rootCmd.Flags().String("proxy-pinnedpubkey", "", "FILE/HASHES public key to verify proxy with")
	rootCmd.Flags().String("proxy-service-name", "", "SPNEGO proxy service name")
	rootCmd.Flags().String("proxy-ssl-allow-beast", "", "Allow this security flaw for HTTPS proxy")
	rootCmd.Flags().Bool("proxy-ssl-auto-client-cert", false, "Auto client certificate for proxy")
	rootCmd.Flags().String("proxy-tls13-ciphers", "", "TLS 1.3 proxy cipher suites")
	rootCmd.Flags().String("proxy-tlsauthtype", "", "TLS authentication type for HTTPS proxy")
	rootCmd.Flags().String("proxy-tlspassword", "", "TLS password for HTTPS proxy")
	rootCmd.Flags().String("proxy-tlsuser", "", "TLS username for HTTPS proxy")
	rootCmd.Flags().Bool("proxy-tlsv1", false, "TLSv1 for HTTPS proxy")
	rootCmd.Flags().StringP("proxy-user", "U", "", "Proxy user and password")
	rootCmd.Flags().String("proxy1.0", "", "Use HTTP/1.0 proxy on given port")
	rootCmd.Flags().BoolP("proxytunnel", "p", false, "HTTP proxy tunnel (using CONNECT)")
	rootCmd.Flags().String("pubkey", "", "SSH Public key filename")
	rootCmd.Flags().BoolP("quote", "Q", false, "Send command(s) to server before transfer")
	rootCmd.Flags().String("random-file", "", "File for reading random data from")
	rootCmd.Flags().StringP("range", "r", "", "Retrieve only the bytes within RANGE")
	rootCmd.Flags().String("rate", "", "Request rate for serial transfers")
	rootCmd.Flags().Bool("raw", false, "Do HTTP \"raw\"; no transfer decoding")
	rootCmd.Flags().StringP("referer", "e", "", "Referrer URL")
	rootCmd.Flags().StringP("remote-header-name", "J", "", "Use the header-provided filename")
	rootCmd.Flags().BoolP("remote-name", "O", false, "Write output to file named as remote file")
	rootCmd.Flags().String("remote-name-all", "", "Use the remote filename for all URLs")
	rootCmd.Flags().BoolP("remote-time", "R", false, "Set remote file's time on local output")
	rootCmd.Flags().Bool("remove-on-error", false, "Remove output file on errors")
	rootCmd.Flags().StringP("request", "X", "", "Specify request method to use")
	rootCmd.Flags().String("request-target", "", "Specify the target for this request")
	rootCmd.Flags().String("resolve", "", "Resolve host+port to address")
	rootCmd.Flags().String("retry", "", "Retry request if transient problems occur")
	rootCmd.Flags().Bool("retry-all-errors", false, "Retry all errors (with --retry)")
	rootCmd.Flags().String("retry-connrefused", "", "Retry on connection refused (with --retry)")
	rootCmd.Flags().String("retry-delay", "", "Wait time between retries")
	rootCmd.Flags().String("retry-max-time", "", "Retry only within this period")
	rootCmd.Flags().String("sasl-authzid", "", "Identity for SASL PLAIN authentication")
	rootCmd.Flags().Bool("sasl-ir", false, "Initial response in SASL authentication")
	rootCmd.Flags().String("service-name", "", "SPNEGO service name")
	rootCmd.Flags().BoolP("show-error", "S", false, "Show error even when -s is used")
	rootCmd.Flags().BoolP("show-headers", "i", false, "Show response headers in output")
	rootCmd.Flags().String("sigalgs", "", "TLS signature algorithms to use")
	rootCmd.Flags().BoolP("silent", "s", false, "Silent mode")
	rootCmd.Flags().Bool("skip-existing", false, "Skip download if local file already exists")
	rootCmd.Flags().String("socks4", "", "SOCKS4 proxy on given host + port")
	rootCmd.Flags().String("socks4a", "", "SOCKS4a proxy on given host + port")
	rootCmd.Flags().String("socks5", "", "SOCKS5 proxy on given host + port")
	rootCmd.Flags().Bool("socks5-basic", false, "Username/password auth for SOCKS5 proxies")
	rootCmd.Flags().String("socks5-gssapi", "", "Enable GSS-API auth for SOCKS5 proxies")
	rootCmd.Flags().String("socks5-gssapi-nec", "", "Compatibility with NEC SOCKS5 server")
	rootCmd.Flags().String("socks5-gssapi-service", "", "SOCKS5 proxy service name for GSS-API")
	rootCmd.Flags().String("socks5-hostname", "", "SOCKS5 proxy, pass hostname to proxy")
	rootCmd.Flags().StringP("speed-limit", "Y", "", "Stop transfers slower than this")
	rootCmd.Flags().StringP("speed-time", "y", "", "Trigger 'speed-limit' abort after this time")
	rootCmd.Flags().Bool("ssl", false, "Try enabling TLS")
	rootCmd.Flags().String("ssl-allow-beast", "", "Allow security flaw to improve interop")
	rootCmd.Flags().Bool("ssl-auto-client-cert", false, "Use auto client certificate (Schannel)")
	rootCmd.Flags().String("ssl-no-revoke", "", "Disable cert revocation checks (Schannel)")
	rootCmd.Flags().Bool("ssl-reqd", false, "Require SSL/TLS")
	rootCmd.Flags().String("ssl-revoke-best-effort", "", "Ignore missing cert CRL dist points")
	rootCmd.Flags().String("ssl-sessions", "", "Load/save SSL session tickets from/to this file")
	rootCmd.Flags().BoolP("sslv2", "2", false, "SSLv2")
	rootCmd.Flags().BoolP("sslv3", "3", false, "SSLv3")
	rootCmd.Flags().Bool("stderr", false, "Where to redirect stderr")
	rootCmd.Flags().String("styled-output", "", "Enable styled output for HTTP headers")
	rootCmd.Flags().String("suppress-connect-headers", "", "Suppress proxy CONNECT response headers")
	rootCmd.Flags().Bool("tcp-fastopen", false, "Use TCP Fast Open")
	rootCmd.Flags().Bool("tcp-nodelay", false, "Set TCP_NODELAY")
	rootCmd.Flags().StringP("telnet-option", "t", "", "Set telnet option")
	rootCmd.Flags().String("tftp-blksize", "", "Set TFTP BLKSIZE option")
	rootCmd.Flags().String("tftp-no-options", "", "Do not send any TFTP options")
	rootCmd.Flags().StringP("time-cond", "z", "", "Transfer based on a time condition")
	rootCmd.Flags().Bool("tls-earlydata", false, "Allow use of TLSv1.3 early data (0RTT)")
	rootCmd.Flags().String("tls-max", "", "Maximum allowed TLS version")
	rootCmd.Flags().String("tls13-ciphers", "", "TLS 1.3 cipher suites to use")
	rootCmd.Flags().String("tlsauthtype", "", "TLS authentication type")
	rootCmd.Flags().Bool("tlspassword", false, "TLS password")
	rootCmd.Flags().String("tlsuser", "", "TLS username")
	rootCmd.Flags().BoolP("tlsv1", "1", false, "TLSv1.0 or greater")
	rootCmd.Flags().Bool("tlsv1.0", false, "TLSv1.0 or greater")
	rootCmd.Flags().Bool("tlsv1.1", false, "TLSv1.1 or greater")
	rootCmd.Flags().Bool("tlsv1.2", false, "TLSv1.2 or greater")
	rootCmd.Flags().Bool("tlsv1.3", false, "TLSv1.3 or greater")
	rootCmd.Flags().Bool("tr-encoding", false, "Request compressed transfer encoding")
	rootCmd.Flags().String("trace", "", "Write a debug trace to FILE")
	rootCmd.Flags().String("trace-ascii", "", "Like --trace, but without hex output")
	rootCmd.Flags().StringSlice("trace-config", nil, "Details to log in trace/verbose output")
	rootCmd.Flags().Bool("trace-ids", false, "Transfer + connection ids in verbose output")
	rootCmd.Flags().Bool("trace-time", false, "Add time stamps to trace/verbose output")
	rootCmd.Flags().String("unix-socket", "", "Connect through this Unix domain socket")
	rootCmd.Flags().StringP("upload-file", "T", "", "Transfer local FILE to destination")
	rootCmd.Flags().String("upload-flags", "", "IMAP upload behavior")
	rootCmd.Flags().String("url", "", "URL(s) to work with")
	rootCmd.Flags().String("url-query", "", "Add a URL query part")
	rootCmd.Flags().BoolP("use-ascii", "B", false, "Use ASCII/text transfer")
	rootCmd.Flags().StringP("user", "u", "", "Server user and password")
	rootCmd.Flags().StringP("user-agent", "A", "", "Send User-Agent <name> to server")
	rootCmd.Flags().String("variable", "", "Set variable")
	rootCmd.Flags().BoolP("verbose", "v", false, "Make the operation more talkative")
	rootCmd.Flags().BoolP("version", "V", false, "Show version number and quit")
	rootCmd.Flags().String("vlan-priority", "", "Set VLAN priority")
	rootCmd.Flags().StringP("write-out", "w", "", "Output FORMAT after completion")
	rootCmd.Flags().Bool("xattr", false, "Store metadata in extended file attributes")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"alt-svc":   carapace.ActionFiles(),
		"cacert":    carapace.ActionFiles(),
		"capath":    carapace.ActionDirectories(),
		"cert":      carapace.ActionFiles(),
		"cert-type": carapace.ActionValues("DER", "PEM", "ENG", "PROV", "P12"),
		"ciphers":   ssh.ActionCiphers().UniqueList(":"),
		"config":    carapace.ActionFiles(),
		"connect-to": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts().Invoke(c).Suffix(":").ToA()
			case 1:
				return net.ActionPorts().Invoke(c).Suffix(":").ToA()
			case 2:
				return net.ActionHosts().Invoke(c).Suffix(":").ToA()
			case 3:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
		"cookie":     carapace.ActionFiles(),
		"cookie-jar": carapace.ActionFiles(),
		"crlfile":    carapace.ActionFiles(),
		"ech": carapace.ActionValuesDescribed(
			"false", "Do not attempt ECH",
			"grease", "Send a GREASE ECH extension",
			"true", "Attempt ECH if possible, but do not fail if ECH is not attempted",
			"hard", "Attempt ECH and fail if that is not possible",
			"ecl:", "A base64 encoded ECHConfigList that is used for ECH",
			"pn:", "A name to use to over-ride the \"public_name\" field of an ECHConfigList",
		),
		"egd-file":     carapace.ActionFiles(),
		"engine":       action.ActionEngines(),
		"etag-compare": carapace.ActionFiles(),
		"etag-save":    carapace.ActionFiles(),
		"ftp-method": carapace.ActionValuesDescribed(
			"multicwd", "curl does a single CWD operation for each path part in the given URL",
			"nocwd", "curl does no CWD at all",
			"singlecwd", "curl does one CWD with the full target directory and then operates on the file \"normally\"",
		),
		"header": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return http.ActionRequestHeaderNames().Invoke(c).Suffix(":").ToA()
			case 1:
				return http.ActionRequestHeaderValues(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
		"hsts":            carapace.ActionFiles(),
		"ip-tos":          carapace.ActionValues("CS0", "CS1", "CS2", "CS3", "CS4", "CS5", "CS6", "CS7", "AF11", "AF12", "AF13", "AF21", "AF22", "AF23", "AF31", "AF32", "AF33", "AF41", "AF42", "AF43", "EF", "VOICE-ADMIT", "ECT1", "ECT0", "CE", "LE", "LOWCOST", "LOWDELAY", "THROUGHPUT", "RELIABILITY", "MINCOST"),
		"key":             carapace.ActionFiles(),
		"key-type":        carapace.ActionValues("DER", "PEM", "ENG"),
		"krb":             carapace.ActionValues("clear", "safe", "confidential", "private"),
		"output":          carapace.ActionFiles(),
		"output-dir":      carapace.ActionDirectories(),
		"proxy-cacert":    carapace.ActionFiles(),
		"proxy-capath":    carapace.ActionDirectories(),
		"proxy-cert":      carapace.ActionFiles(),
		"proxy-cert-type": carapace.ActionValues("DER", "PEM", "ENG"),
		"proxy-ciphers":   ssh.ActionCiphers().UniqueList(":"),
		"proxy-key":       carapace.ActionFiles(),
		"random-file":     carapace.ActionFiles(),
		"ssl-sessions":    carapace.ActionFiles(),
		"trace":           carapace.ActionFiles(),
		"unix-socket":     carapace.ActionFiles(),
		"upload-file":     carapace.ActionFiles(),
		"upload-flags":    carapace.ActionValues("answered", "deleted", "draft", "flagged", "seen").UniqueList(","),
		"user-agent":      http.ActionUserAgents(),
		"vlan-priority":   carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		http.ActionUrls(),
	)
}
