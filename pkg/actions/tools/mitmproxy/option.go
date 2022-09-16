package mitmproxy

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
	"github.com/rsteube/carapace-bin/pkg/actions/net/ssh"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionOptionNames completion option names
//
//	anticomp (Try to convince servers to send us un-compressed data)
//	block_global (Block connections from public IP addresses)
func ActionOptionNames() carapace.Action {
	return carapace.ActionValuesDescribed(
		"add_upstream_certs_to_client_chain", "Add all certificates of the upstream server to the certificate chain",
		"allow_hosts", "Opposite of --ignore-hosts",
		"anticache", "Strip out request headers that might cause the server to return",
		"anticomp", "Try to convince servers to send us un-compressed data",
		"block_global", "Block connections from public IP addresses",
		"block_list", "Block matching requests and return an empty response",
		"block_private", "Block connections from local (private) IP addresses",
		"body_size_limit", "Byte size limit of HTTP request and response bodies",
		"cert_passphrase", "Passphrase for decrypting the private key provided in the --cert option",
		"certs", "SSL certificates of the form \"[domain=]path\"",
		"ciphers_client", "Set supported ciphers for client <-> mitmproxy connections",
		"ciphers_server", "Set supported ciphers for mitmproxy <-> server connections using OpenSSL syntax",
		"client_certs", "Client certificate file or directory",
		"client_replay", "Replay client requests from a saved file",
		"client_replay_concurrency", "Concurrency limit on in-flight client replay requests",
		"command_history", "Persist command history between mitmproxy invocations",
		"confdir", "Location of the default mitmproxy configuration files",
		"connect_addr", "Set the local IP address that mitmproxy should use when connecting to upstream servers",
		"connection_strategy", "Determine when server connections should be established",
		"console_default_contentview", "The default content view mode",
		"console_eventlog_verbosity", "EventLog verbosity",
		"console_flowlist_layout", "Set the flowlist layout",
		"console_focus_follow", "Focus follows new flows",
		"console_layout", "Console layout",
		"console_layout_headers", "Show layout component headers",
		"console_mouse", "Console mouse interaction",
		"console_palette", "Color palette",
		"console_palette_transparent", "Set transparent background for palette",
		"console_strip_trailing_newlines", "Strip trailing newlines from edited request/response bodies",
		"content_view_lines_cutoff", "Flow content view lines limit",
		"dns_listen_host", "Address to bind DNS server to",
		"dns_listen_port", "DNS server service port",
		"dns_mode", "DNS mode",
		"dns_server", "Start a DNS server",
		"export_preserve_original_ip", "make an effort to connect to the same IP as in the original request",
		"http2", "Enable/disable HTTP/2 support",
		"http2_ping_keepalive", "Send a PING frame if an HTTP/2 connection is idle for specified number of seconds",
		"ignore_hosts", "Ignore host and forward all traffic without processing it",
		"intercept", "Intercept filter expression",
		"intercept_active", "Intercept toggle",
		"keep_host_header", "Reverse Proxy: Keep the original host header instead of rewriting it to the reverse proxy target",
		"key_size", "TLS key size for certificates and CA",
		"listen_host", "Address to bind proxy to",
		"listen_port", "Proxy service port",
		"map_local", "Map remote resources to a local file using a pattern of the form \"[/flow-filter]/url-regex/file-or-directory-path\"",
		"map_remote", "Map remote resources to another remote URL using a pattern of the form \"[/flow-filter]/url-regex/replacement\"",
		"mode", "proxy mode",
		"modify_body", "Replacement pattern of the form \"[/flow-filter]/regex/[@]replacement\"",
		"modify_headers", "Header modify pattern of the form \"[/flow-filter]/header-name/[@]header-value\"",
		"normalize_outbound_headers", "Normalize outgoing HTTP/2 header names",
		"onboarding", "Toggle the mitmproxy onboarding app",
		"onboarding_host", "Onboarding app domain",
		"onboarding_port", "Port to serve the onboarding app from",
		"proxy_debug", "Enable debug logs in the proxy core",
		"proxyauth", "Require proxy authentication",
		"rawtcp", "Enable/disable raw TCP connections",
		"readfile_filter", "Read only matching flows",
		"rfile", "Read flows from file",
		"save_stream_file", "Stream flows to file as they arrive",
		"save_stream_filter", "Filter which flows are written to file",
		"scripts", "Execute a script",
		"server", "Start a proxy server",
		"server_replay", "Replay server responses from a saved file",
		"server_replay_ignore_content", "Ignore request content while searching for a saved flow to replay",
		"server_replay_ignore_host", "Ignore request destination host while searching for a saved flow to replay",
		"server_replay_ignore_params", "Request parameters to be ignored while searching for a saved flow to replay",
		"server_replay_ignore_payload_params", "Request payload parameters to be ignored while searching for a saved flow to replay",
		"server_replay_ignore_port", "Ignore request destination port while searching for a saved flow to replay",
		"server_replay_kill_extra", "Kill extra requests during replay",
		"server_replay_nopop", "Don't remove flows from server replay state after use",
		"server_replay_refresh", "Refresh server replay responses by adjusting headers",
		"server_replay_use_headers", "Request headers that need to match while searching for a saved flow to replay",
		"showhost", "Use the Host header to construct URLs for display",
		"ssl_insecure", "Do not verify upstream server SSL/TLS certificates",
		"ssl_verify_upstream_trusted_ca", "Path to a PEM formatted trusted CA certificate",
		"ssl_verify_upstream_trusted_confdir", "Path to a directory of trusted CA certificates for upstream server verification",
		"stickyauth", "Set sticky auth filter",
		"stickycookie", "Set sticky cookie filter. Matched against requests",
		"stream_large_bodies", "Stream data to the client if response body exceeds the given threshold",
		"tcp_hosts", "Generic TCP SSL proxy mode for all hosts that match the pattern",
		"tls_version_client_max", "Set the maximum TLS version for client connections",
		"tls_version_client_min", "Set the minimum TLS version for client connections",
		"tls_version_server_max", "Set the maximum TLS version for server connections",
		"tls_version_server_min", "Set the minimum TLS version for server connections",
		"upstream_auth", "Add HTTP Basic authentication to upstream proxy and reverse proxy requests",
		"upstream_cert", "Connect to upstream server to look up certificate details",
		"validate_inbound_headers", "Make sure that incoming HTTP requests are not malformed",
		"view_filter", "Limit the view to matching flows",
		"view_order", "Flow sort order",
		"view_order_reversed", "Reverse the sorting order",
		"websocket", "Enable/disable WebSocket support",
	)
}

// ActionOptionValues completes option validate_inbound_headers
//
//	auto
//	css
func ActionOptionValues(option string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		actionBool := carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
		actionTlsVersions := carapace.ActionValues("UNBOUNDED", "SSL3", "TLS1", "TLS1_1", "TLS1_2", "TLS1_3")
		return map[string]carapace.Action{
			"add_upstream_certs_to_client_chain": actionBool,
			"allow_hosts":                        carapace.ActionValues(),
			"anticache":                          actionBool,
			"anticomp":                           actionBool,
			"block_global":                       actionBool,
			"block_list": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if len(c.CallbackValue) > 0 {
					return carapace.ActionMultiParts(c.CallbackValue[:1], func(c carapace.Context) carapace.Action {
						switch len(c.Parts) {
						case 2:
							return http.ActionStatusCodes()
						default:
							return carapace.ActionValues()
						}
					})
				}
				return carapace.ActionValues()
			}),
			"block_private":   actionBool,
			"body_size_limit": carapace.ActionValues(),
			"cert_passphrase": carapace.ActionValues(),
			"certs": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				if len(c.Parts) > 0 {
					return carapace.ActionFiles()
				}
				return carapace.ActionValues()
			}),
			"ciphers_client":            ssh.ActionCiphers(),
			"ciphers_server":            ssh.ActionCiphers(),
			"client_certs":              carapace.ActionFiles(),
			"client_replay":             carapace.ActionFiles(),
			"client_replay_concurrency": carapace.ActionValues(),
			"command_history":           actionBool,
			"confdir":                   carapace.ActionDirectories(),
			"connect_addr":              carapace.ActionValues(),
			"connection_strategy":       carapace.ActionValues("eager", "lazy"),
			"console_default_contentview": carapace.ActionValues("auto", "raw", "hex", "graphql", "json", "xml/html", "wbxml", "javascript",
				"css", "url-encoded", "multipart form", "image", "query", "protocol buffer", "msgpack", "grpc/protocol buffer"),
			"console_eventlog_verbosity":      carapace.ActionValues("error", "warn", "info", "alert", "debug").StyleF(style.ForLogLevel),
			"console_flowlist_layout":         carapace.ActionValues("default", "list", "table"),
			"console_focus_follow":            actionBool,
			"console_layout":                  ActionConsoleLayouts(),
			"console_layout_headers":          actionBool,
			"console_mouse":                   actionBool,
			"console_palette":                 carapace.ActionValues("dark", "light", "lowdark", "lowlight", "solarized_dark", "solarized_light"),
			"console_palette_transparent":     actionBool,
			"console_strip_trailing_newlines": actionBool,
			"content_view_lines_cutoff":       carapace.ActionValues(),
			"dns_listen_host":                 carapace.ActionValues(),
			"dns_listen_port":                 net.ActionPorts(),
			"dns_mode": carapace.Batch(
				carapace.ActionValuesDescribed(
					"regular", "requests will be resolved using the local resolver",
					"transparent", "transparent mode",
				),
				carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValuesDescribed("reverse", "forward queries to another DNS server transparent").Invoke(c).Suffix(":").ToA()
					case 1:
						return net.ActionIpv4Addresses()
					case 2:
						return net.ActionPorts()
					default:
						return carapace.ActionValues()

					}
				}),
			).ToA(),
			"dns_server":                   actionBool,
			"export_preserve_original_ip":  actionBool,
			"http2":                        actionBool,
			"http2_ping_keepalive":         carapace.ActionValues(),
			"ignore_hosts":                 carapace.ActionValues(),
			"intercept":                    carapace.ActionValues(), // TODO completion?
			"intercept_active":             actionBool,
			"keep_host_header":             actionBool,
			"key_size":                     carapace.ActionValues(),
			"listen_host":                  carapace.ActionValues(),
			"listen_port":                  net.ActionPorts(),
			"map_local":                    carapace.ActionValues(), // TODO completion
			"map_remote":                   carapace.ActionValues(),
			"mode":                         ActionModes(),
			"modify_body":                  ActionModifyBodyPattern(),
			"modify_headers":               ActionModifyHeaderPattern(),
			"normalize_outbound_headers":   actionBool,
			"onboarding":                   actionBool,
			"onboarding_host":              carapace.ActionValues(),
			"onboarding_port":              net.ActionPorts(),
			"proxy_debug":                  actionBool,
			"proxyauth":                    carapace.ActionValues(),
			"rawtcp":                       actionBool,
			"readfile_filter":              actionOptionFlowFilters(),
			"rfile":                        carapace.ActionFiles(),
			"save_stream_file":             ActionAppendableFiles(),
			"save_stream_filter":           actionOptionFlowFilters(),
			"scripts":                      carapace.ActionFiles(),
			"server":                       actionBool,
			"server_replay":                carapace.ActionFiles(),
			"server_replay_ignore_content": actionBool,
			"server_replay_ignore_host":    actionBool,
			"server_replay_ignore_params":  carapace.ActionValues(),
			"server_replay_ignore_payload_params": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				return http.ActionMediaTypes().Invoke(c).ToMultiPartsA("/")
			}),
			"server_replay_ignore_port":           actionBool,
			"server_replay_kill_extra":            actionBool,
			"server_replay_nopop":                 actionBool,
			"server_replay_refresh":               actionBool,
			"server_replay_use_headers":           http.ActionRequestHeaderNames(),
			"showhost":                            actionBool,
			"ssl_insecure":                        actionBool,
			"ssl_verify_upstream_trusted_ca":      carapace.ActionFiles(),
			"ssl_verify_upstream_trusted_confdir": carapace.ActionDirectories(),
			"stickyauth":                          carapace.ActionValues(), // TODO completion?
			"stickycookie":                        carapace.ActionValues(), // TODO completion?
			"stream_large_bodies":                 carapace.ActionValues(),
			"tcp_hosts":                           carapace.ActionValues(),
			"tls_version_client_max":              actionTlsVersions,
			"tls_version_client_min":              actionTlsVersions,
			"tls_version_server_max":              actionTlsVersions,
			"tls_version_server_min":              actionTlsVersions,
			"upstream_auth":                       carapace.ActionValues(),
			"upstream_cert":                       actionBool,
			"validate_inbound_headers":            actionBool,
			"view_filter":                         carapace.ActionValues(), // TODO completion?
			"view_order":                          carapace.ActionValues("time", "method", "url", "size"),
			"view_order_reversed":                 actionBool,
			"websocket":                           actionBool,
		}[option]
	})
}

func actionOptionFiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(c.CallbackValue, "@") {
			c.CallbackValue = strings.TrimPrefix(c.CallbackValue, "@")
			return carapace.ActionFiles().Invoke(c).Prefix("@").ToA()
		}
		return carapace.ActionValues()
	})
}

func actionOptionFlowFilters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if index := strings.LastIndex(c.CallbackValue, "~"); index != -1 {
			return ActionFlowFilters().Invoke(c).Prefix(c.CallbackValue[:index+1]).ToA()
		}
		return carapace.ActionValues()
	})

}
