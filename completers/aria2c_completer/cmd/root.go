package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aria2c",
	Short: "The ultra fast download utility",
	Long:  "https://aria2.github.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("all-proxy", "", "Use a proxy server for all protocols")
	rootCmd.Flags().String("all-proxy-passwd", "", "Set password for --all-proxy")
	rootCmd.Flags().String("all-proxy-user", "", "Set user for --all-proxy")
	rootCmd.Flags().Bool("allow-overwrite", false, "Restart download from scratch if the corresponding control file doesn't exist")
	rootCmd.Flags().Bool("allow-piece-length-change", false, "If false is given, aria2 aborts download when a piece length is different from one in a control file")
	rootCmd.Flags().Bool("always-resume", false, "Always resume download")
	rootCmd.Flags().Bool("async-dns", false, "Enable asynchronous DNS")
	rootCmd.Flags().String("async-dns-server", "", "Comma separated list of DNS server address used in asynchronous DNS resolver")
	rootCmd.Flags().Bool("auto-file-renaming", false, "Rename file name if the same file already exists")
	rootCmd.Flags().String("auto-save-interval", "", "Save a control file(*.aria2) every SEC seconds")
	rootCmd.Flags().Bool("bt-detach-seed-only", false, "Exclude seed only downloads when counting concurrent active downloads")
	rootCmd.Flags().Bool("bt-enable-hook-after-hash-check", false, "Allow hook command invocation after hash check (see -V option) in BitTorrent download")
	rootCmd.Flags().Bool("bt-enable-lpd", false, "Enable Local Peer Discovery")
	rootCmd.Flags().String("bt-exclude-tracker", "", "Comma separated list of BitTorrent tracker's announce URI to remove")
	rootCmd.Flags().String("bt-external-ip", "", "Specify the external IP address to use in BitTorrent download and DHT")
	rootCmd.Flags().Bool("bt-force-encryption", false, "Requires BitTorrent message payload encryption with arc4")
	rootCmd.Flags().Bool("bt-hash-check-seed", false, "If true is given, after hash check using --check-integrity option and file is complete, continue to seed file")
	rootCmd.Flags().Bool("bt-load-saved-metadata", false, "Before getting torrent metadata from DHT when downloading with magnet link, first try to read file saved by --bt-save-metadata option")
	rootCmd.Flags().String("bt-lpd-interface", "", "Use given interface for Local Peer Discovery")
	rootCmd.Flags().String("bt-max-open-files", "", "Specify maximum number of files to open in multi-file BitTorrent/Metalink downloads globally")
	rootCmd.Flags().String("bt-max-peers", "", "Specify the maximum number of peers per torrent")
	rootCmd.Flags().Bool("bt-metadata-only", false, "Download metadata only")
	rootCmd.Flags().String("bt-min-crypto-level", "", "Set minimum level of encryption method")
	rootCmd.Flags().String("bt-prioritize-piece", "", "Try to download first and last pieces of each file first")
	rootCmd.Flags().Bool("bt-remove-unselected-file", false, "Removes the unselected files when download is completed in BitTorrent")
	rootCmd.Flags().String("bt-request-peer-speed-limit", "", "If the whole download speed of every torrent is lower than SPEED, aria2 temporarily increases the number of peers to try for more download speed")
	rootCmd.Flags().Bool("bt-require-crypto", false, "If true is given, aria2 doesn't accept and establish connection with legacy BitTorrent handshake")
	rootCmd.Flags().Bool("bt-save-metadata", false, "Save metadata as .torrent file")
	rootCmd.Flags().Bool("bt-seed-unverified", false, "Seed previously downloaded files without verifying piece hashes")
	rootCmd.Flags().String("bt-stop-timeout", "", "Stop BitTorrent download if download speed is 0 in consecutive SEC seconds")
	rootCmd.Flags().String("bt-tracker", "", "Comma separated list of additional BitTorrent tracker's announce URI")
	rootCmd.Flags().String("bt-tracker-connect-timeout", "", "Set the connect timeout in seconds to establish connection to tracker")
	rootCmd.Flags().String("bt-tracker-interval", "", "Set the interval in seconds between tracker requests")
	rootCmd.Flags().String("bt-tracker-timeout", "", "Set timeout in seconds")
	rootCmd.Flags().String("ca-certificate", "", "Use the certificate authorities in FILE to verify the peers")
	rootCmd.Flags().String("certificate", "", "Use the client certificate in FILE")
	rootCmd.Flags().Bool("check-certificate", false, "Verify the peer using certificates specified in --ca-certificate option")
	rootCmd.Flags().BoolP("check-integrity", "V", false, "Check file integrity by validating piece hashes or a hash of entire file")
	rootCmd.Flags().String("checksum", "", "Set checksum")
	rootCmd.Flags().Bool("conditional-get", false, "Download file only when the local file is older than remote file")
	rootCmd.Flags().String("conf-path", "", "Change the configuration file path to PATH")
	rootCmd.Flags().String("connect-timeout", "", "Set the connect timeout in seconds to establish connection to HTTP/FTP/proxy server")
	rootCmd.Flags().String("console-log-level", "", "Set log level to output to console")
	rootCmd.Flags().Bool("content-disposition-default-utf8", false, "Handle quoted string in Content-Disposition header as UTF-8 instead of ISO-8859-1, for example, the filename parameter, but not the extended version filename*")
	rootCmd.Flags().BoolP("continue", "c", false, "Continue downloading a partially downloaded file")
	rootCmd.Flags().BoolP("daemon", "D", false, "Run as daemon")
	rootCmd.Flags().Bool("deferred-input", false, "If true is given, aria2 does not read all URIs and options from file specified by -i option at startup, but it reads one by one when it needs later")
	rootCmd.Flags().String("dht-entry-point", "", "Set host and port as an entry point to IPv4 DHT network")
	rootCmd.Flags().String("dht-entry-point6", "", "Set host and port as an entry point to IPv6 DHT network")
	rootCmd.Flags().String("dht-file-path", "", "Change the IPv4 DHT routing table file to PATH")
	rootCmd.Flags().String("dht-file-path6", "", "Change the IPv6 DHT routing table file to PATH")
	rootCmd.Flags().String("dht-listen-addr6", "", "Specify address to bind socket for IPv6 DHT")
	rootCmd.Flags().String("dht-listen-port", "", "Set UDP listening port used by DHT(IPv4, IPv6) and UDP tracker")
	rootCmd.Flags().String("dht-message-timeout", "", "Set timeout in seconds")
	rootCmd.Flags().StringP("dir", "d", "", "The directory to store the downloaded file")
	rootCmd.Flags().Bool("disable-ipv6", false, "Disable IPv6")
	rootCmd.Flags().String("disk-cache", "", "Enable disk cache")
	rootCmd.Flags().String("download-result", "", "This option changes the way \"Download Results\" is formatted")
	rootCmd.Flags().Bool("dry-run", false, "If true is given, aria2 just checks whether the remote file is available and doesn't download data")
	rootCmd.Flags().String("dscp", "", "Set DSCP value in outgoing IP packets of BitTorrent traffic for QoS")
	rootCmd.Flags().Bool("enable-async-dns6", false, "Enable IPv6 name resolution in asynchronous DNS resolver")
	rootCmd.Flags().Bool("enable-color", false, "Enable color output for a terminal")
	rootCmd.Flags().Bool("enable-dht", false, "Enable IPv4 DHT functionality")
	rootCmd.Flags().Bool("enable-dht6", false, "Enable IPv6 DHT functionality")
	rootCmd.Flags().Bool("enable-http-keep-alive", false, "Enable HTTP/1.1 persistent connection")
	rootCmd.Flags().Bool("enable-http-pipelining", false, "Enable HTTP/1.1 pipelining")
	rootCmd.Flags().Bool("enable-mmap", false, "Map files into memory")
	rootCmd.Flags().Bool("enable-peer-exchange", false, "Enable Peer Exchange extension")
	rootCmd.Flags().Bool("enable-rpc", false, "Enable JSON-RPC/XML-RPC server")
	rootCmd.Flags().String("event-poll", "", "Specify the method for polling events")
	rootCmd.Flags().String("file-allocation", "", "Specify file allocation method")
	rootCmd.Flags().String("follow-metalink", "", "Controls how Metalink files are handled after download")
	rootCmd.Flags().String("follow-torrent", "", "Controls how Torrent files are handled after download")
	rootCmd.Flags().Bool("force-save", false, "Save download with --save-session option even if the download is completed or removed")
	rootCmd.Flags().BoolP("force-sequential", "Z", false, "Fetch URIs in the command-line sequentially and download each URI in a separate session")
	rootCmd.Flags().String("ftp-passwd", "", "Set FTP password")
	rootCmd.Flags().BoolP("ftp-pasv", "p", false, "Use the passive mode in FTP")
	rootCmd.Flags().String("ftp-proxy", "", "Use a proxy server for FTP")
	rootCmd.Flags().String("ftp-proxy-passwd", "", "Set password for --ftp-proxy")
	rootCmd.Flags().String("ftp-proxy-user", "", "Set user for --ftp-proxy")
	rootCmd.Flags().Bool("ftp-reuse-connection", false, "Reuse connection in FTP")
	rootCmd.Flags().String("ftp-type", "", "Set FTP transfer type")
	rootCmd.Flags().String("ftp-user", "", "Set FTP user")
	rootCmd.Flags().String("gid", "", "Set GID manually")
	rootCmd.Flags().Bool("hash-check-only", false, "If true is given, after hash check using --check-integrity option, abort download whether or not download is complete")
	rootCmd.Flags().String("header", "", "Append HEADER to HTTP request header")
	rootCmd.Flags().StringP("help", "h", "", "Print usage and exit")
	rootCmd.Flags().Bool("http-accept-gzip", false, "Send 'Accept-Encoding: deflate, gzip' request header and inflate response if remote server responds with 'Content-Encoding: gzip' or 'Content-Encoding: deflate'")
	rootCmd.Flags().Bool("http-auth-challenge", false, "Send HTTP authorization header only when it is requested by the server")
	rootCmd.Flags().Bool("http-no-cache", false, "Send Cache-Control: no-cache and Pragma: no-cache header to avoid cached content")
	rootCmd.Flags().String("http-passwd", "", "Set HTTP password")
	rootCmd.Flags().String("http-proxy", "", "Use a proxy server for HTTP")
	rootCmd.Flags().String("http-proxy-passwd", "", "Set password for --http-proxy")
	rootCmd.Flags().String("http-proxy-user", "", "Set user for --http-proxy")
	rootCmd.Flags().String("http-user", "", "Set HTTP user")
	rootCmd.Flags().String("https-proxy", "", "Use a proxy server for HTTPS")
	rootCmd.Flags().String("https-proxy-passwd", "", "Set password for --https-proxy")
	rootCmd.Flags().String("https-proxy-user", "", "Set user for --https-proxy")
	rootCmd.Flags().Bool("human-readable", false, "Print sizes and speed in human readable format (e.g., 1.2Ki, 3.4Mi) in the console readout")
	rootCmd.Flags().StringSliceP("index-out", "O", []string{}, "Set file path for file with index=INDEX")
	rootCmd.Flags().StringP("input-file", "i", "", "Downloads URIs found in FILE")
	rootCmd.Flags().String("interface", "", "Bind sockets to given interface")
	rootCmd.Flags().Bool("keep-unfinished-download-result", false, "Keep unfinished download results even if doing so exceeds --max-download-result")
	rootCmd.Flags().String("listen-port", "", "Set TCP port number for BitTorrent downloads")
	rootCmd.Flags().String("load-cookies", "", "Load Cookies from FILE using the Firefox3 format and Mozilla/Firefox(1.x/2.x)/Netscape format")
	rootCmd.Flags().StringP("log", "l", "", "The file name of the log file")
	rootCmd.Flags().String("log-level", "", "Set log level to output to file specified using --log option")
	rootCmd.Flags().String("lowest-speed-limit", "", "Close connection if download speed is lower than or equal to this value(bytes per sec)")
	rootCmd.Flags().StringP("max-concurrent-downloads", "j", "", "Set maximum number of parallel downloads for every static (HTTP/FTP) URL, torrent and metalink")
	rootCmd.Flags().StringP("max-connection-per-server", "x", "", "The maximum number of connections to one server for each download")
	rootCmd.Flags().String("max-download-limit", "", "Set max download speed per each download in bytes/sec")
	rootCmd.Flags().String("max-download-result", "", "Set maximum number of download result kept in memory")
	rootCmd.Flags().String("max-file-not-found", "", "If aria2 receives `file not found' status from the remote HTTP/FTP servers NUM times without getting a single byte, then force the download to fail")
	rootCmd.Flags().String("max-mmap-limit", "", "Set the maximum file size to enable mmap")
	rootCmd.Flags().String("max-overall-download-limit", "", "Set max overall download speed in bytes/sec")
	rootCmd.Flags().String("max-overall-upload-limit", "", "Set max overall upload speed in bytes/sec")
	rootCmd.Flags().String("max-resume-failure-tries", "", "When used with --always-resume=false, aria2 downloads file from scratch when aria2 detects N number of URIs that does not support resume")
	rootCmd.Flags().StringP("max-tries", "m", "", "Set number of tries")
	rootCmd.Flags().StringP("max-upload-limit", "u", "", "Set max upload speed per each torrent in bytes/sec")
	rootCmd.Flags().String("metalink-base-uri", "", "Specify base URI to resolve relative URI in metalink:url and metalink:metaurl element in a metalink file stored in local disk")
	rootCmd.Flags().Bool("metalink-enable-unique-protocol", false, "If true is given and several protocols are available for a mirror in a metalink file, aria2 uses one of them")
	rootCmd.Flags().StringP("metalink-file", "M", "", "The file path to the .meta4 and .metalink file")
	rootCmd.Flags().String("metalink-language", "", "The language of the file to download")
	rootCmd.Flags().String("metalink-location", "", "The location of the preferred server")
	rootCmd.Flags().String("metalink-os", "", "The operating system of the file to download")
	rootCmd.Flags().String("metalink-preferred-protocol", "", "Specify preferred protocol")
	rootCmd.Flags().String("metalink-version", "", "The version of the file to download")
	rootCmd.Flags().StringP("min-split-size", "k", "", "aria2 does not split less than 2*SIZE byte range")
	rootCmd.Flags().String("min-tls-version", "", "Specify minimum SSL/TLS version to enable")
	rootCmd.Flags().String("multiple-interface", "", "Comma separated list of interfaces to bind sockets to")
	rootCmd.Flags().String("netrc-path", "", "Specify the path to the netrc file")
	rootCmd.Flags().Bool("no-conf", false, "Disable loading aria2.conf file")
	rootCmd.Flags().String("no-file-allocation-limit", "", "No file allocation is made for files whose size is smaller than SIZE")
	rootCmd.Flags().BoolP("no-netrc", "n", false, "Disables netrc support")
	rootCmd.Flags().String("no-proxy", "", "Specify comma separated hostnames, domains or network address with or without CIDR block where proxy should not be used")
	rootCmd.Flags().Bool("no-want-digest-header", false, "Whether to disable Want-Digest header when doing requests")
	rootCmd.Flags().String("on-bt-download-complete", "", "For BitTorrent, a command specified in --on-download-complete is called after download completed and seeding is over")
	rootCmd.Flags().String("on-download-complete", "", "Set the command to be executed after download completed")
	rootCmd.Flags().String("on-download-error", "", "Set the command to be executed after download aborted due to error")
	rootCmd.Flags().String("on-download-pause", "", "Set the command to be executed after download was paused")
	rootCmd.Flags().String("on-download-start", "", "Set the command to be executed after download got started")
	rootCmd.Flags().String("on-download-stop", "", "Set the command to be executed after download stopped")
	rootCmd.Flags().Bool("optimize-concurrent-downloads", false, "Optimizes the number of concurrent downloads according to the bandwidth available")
	rootCmd.Flags().StringP("out", "o", "", "The file name of the downloaded file")
	rootCmd.Flags().BoolP("parameterized-uri", "P", false, "Enable parameterized URI support")
	rootCmd.Flags().Bool("pause", false, "Pause download after added")
	rootCmd.Flags().Bool("pause-metadata", false, "Pause downloads created as a result of metadata download")
	rootCmd.Flags().String("peer-agent", "", "Set client reported during Extended torrent handshakes")
	rootCmd.Flags().String("peer-id-prefix", "", "Specify the prefix of peer ID")
	rootCmd.Flags().String("piece-length", "", "Set a piece length for HTTP/FTP downloads")
	rootCmd.Flags().String("private-key", "", "Use the private key in FILE")
	rootCmd.Flags().String("proxy-method", "", "Set the method to use in proxy request")
	rootCmd.Flags().BoolP("quiet", "q", false, "Make aria2 quiet(no console output)")
	rootCmd.Flags().Bool("realtime-chunk-checksum", false, "Validate chunk of data by calculating checksum while downloading a file if chunk checksums are provided")
	rootCmd.Flags().String("referer", "", "Set an http referrrer (Referer)")
	rootCmd.Flags().BoolP("remote-time", "R", false, "Retrieve timestamp of the remote file from the remote HTTP/FTP server and if it is available, apply it to the local file")
	rootCmd.Flags().Bool("remove-control-file", false, "Remove control file before download")
	rootCmd.Flags().String("retry-wait", "", "Set the seconds to wait between retries")
	rootCmd.Flags().Bool("reuse-uri", false, "Reuse already used URIs if no unused URIs are left")
	rootCmd.Flags().String("rlimit-nofile", "", "Set the soft limit of open file descriptors")
	rootCmd.Flags().Bool("rpc-allow-origin-all", false, "Add Access-Control-Allow-Origin header field with value '*' to the RPC response")
	rootCmd.Flags().String("rpc-certificate", "", "Use the certificate in FILE for RPC server")
	rootCmd.Flags().Bool("rpc-listen-all", false, "Listen incoming JSON-RPC/XML-RPC requests on all network interfaces")
	rootCmd.Flags().String("rpc-listen-port", "", "Specify a port number for JSON-RPC/XML-RPC server to listen to")
	rootCmd.Flags().String("rpc-max-request-size", "", "Set max size of JSON-RPC/XML-RPC request")
	rootCmd.Flags().String("rpc-passwd", "", "Set JSON-RPC/XML-RPC password")
	rootCmd.Flags().String("rpc-private-key", "", "Use the private key in FILE for RPC server")
	rootCmd.Flags().Bool("rpc-save-upload-metadata", false, "Save the uploaded torrent or metalink metadata in the directory specified by --dir option")
	rootCmd.Flags().String("rpc-secret", "", "Set RPC secret authorization token")
	rootCmd.Flags().Bool("rpc-secure", false, "RPC transport will be encrypted by SSL/TLS")
	rootCmd.Flags().String("rpc-user", "", "Set JSON-RPC/XML-RPC user")
	rootCmd.Flags().String("save-cookies", "", "Save Cookies to FILE in Mozilla/Firefox(1.x/2.x)/ Netscape format")
	rootCmd.Flags().Bool("save-not-found", false, "Save download with --save-session option even if the file was not found on the server")
	rootCmd.Flags().String("save-session", "", "Save error/unfinished downloads to FILE on exit")
	rootCmd.Flags().String("save-session-interval", "", "Save error/unfinished downloads to a file specified by --save-session option every SEC seconds")
	rootCmd.Flags().String("seed-ratio", "", "Specify share ratio")
	rootCmd.Flags().String("seed-time", "", "Specify seeding time in (fractional) minutes")
	rootCmd.Flags().String("select-file", "", "Set file to download by specifying its index")
	rootCmd.Flags().String("server-stat-if", "", "Specify the filename to load performance profile of the servers")
	rootCmd.Flags().String("server-stat-of", "", "Specify the filename to which performance profile of the servers is saved")
	rootCmd.Flags().String("server-stat-timeout", "", "Specifies timeout in seconds to invalidate performance profile of the servers since the last contact to them")
	rootCmd.Flags().Bool("show-console-readout", false, "Show console readout")
	rootCmd.Flags().BoolP("show-files", "S", false, "Print file listing of .torrent, .meta4 and .metalink file and exit")
	rootCmd.Flags().String("socket-recv-buffer-size", "", "Set the maximum socket receive buffer in bytes")
	rootCmd.Flags().StringP("split", "s", "", "Download a file using N connections")
	rootCmd.Flags().String("ssh-host-key-md", "", "Set checksum for SSH host public key")
	rootCmd.Flags().Bool("stderr", false, "Redirect all console output that would be otherwise printed in stdout to stderr")
	rootCmd.Flags().String("stop", "", "Stop application after SEC seconds has passed")
	rootCmd.Flags().String("stop-with-process", "", "Stop application when process PID is not running")
	rootCmd.Flags().String("stream-piece-selector", "", "Specify piece selection algorithm used in HTTP/FTP download")
	rootCmd.Flags().String("summary-interval", "", "Set interval to output download progress summary")
	rootCmd.Flags().StringP("timeout", "t", "", "Set timeout in seconds")
	rootCmd.Flags().StringP("torrent-file", "T", "", "The path to the .torrent file")
	rootCmd.Flags().Bool("truncate-console-readout", false, "Truncate console readout to fit in a single line")
	rootCmd.Flags().String("uri-selector", "", "Specify URI selection algorithm")
	rootCmd.Flags().Bool("use-head", false, "Use HEAD method for the first request to the HTTP server")
	rootCmd.Flags().StringP("user-agent", "U", "", "Set user agent for http(s) downloads")
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bt-min-crypto-level":         carapace.ActionValues("plain", "arc4"),
		"ca-certificate":              carapace.ActionFiles(),
		"certificate":                 carapace.ActionFiles(),
		"conf-path":                   carapace.ActionFiles(),
		"console-log-level":           carapace.ActionValues("debug", "info", "notice", "warn", "error").StyleF(style.ForLogLevel),
		"dht-file-path":               carapace.ActionFiles(),
		"dht-file-path6":              carapace.ActionFiles(),
		"dir":                         carapace.ActionDirectories(),
		"download-result":             carapace.ActionValues("default", "full", "hide"),
		"event-poll":                  carapace.ActionValues("epll", "poll", "select"),
		"file-allocation":             carapace.ActionValues("none", "prealloc", "trunc", "falloc"),
		"follow-metalink":             carapace.ActionValues("true", "mem", "false").StyleF(style.ForKeyword),
		"follow-torrent":              carapace.ActionValues("true", "mem", "false").StyleF(style.ForKeyword),
		"ftp-type":                    carapace.ActionValues("binary", "ascii"),
		"help":                        carapace.ActionValues("#basic", "#advanced", "#http", "#https", "#ftp", "#metalink", "#bittorrent", "#cookie", "#hook", "#file", "#rpc", "#checksum", "#experimental", "#deprecated", "#help", "#all"),
		"input-file":                  carapace.ActionFiles(),
		"load-cookies":                carapace.ActionFiles(),
		"log":                         carapace.ActionFiles(),
		"log-level":                   carapace.ActionValues("debug", "info", "notice", "warn", "error").StyleF(style.ForLogLevel),
		"metalink-file":               carapace.ActionFiles(),
		"metalink-preferred-protocol": carapace.ActionValues("http", "https", "ftp", "none"),
		"min-tls-version":             carapace.ActionValues("TLSv1.1", "TLSv1.2", "TLSv1.3"),
		"netrc-path":                  carapace.ActionFiles(),
		"on-bt-download-complete":     carapace.ActionFiles(),
		"on-download-complete":        carapace.ActionFiles(),
		"on-download-error":           carapace.ActionFiles(),
		"on-download-pause":           carapace.ActionFiles(),
		"on-download-start":           carapace.ActionFiles(),
		"on-download-stop":            carapace.ActionFiles(),
		"out":                         carapace.ActionFiles(),
		"private-key":                 carapace.ActionFiles(),
		"proxy-method":                carapace.ActionValues("get", "tunnel"),
		"rpc-certificate":             carapace.ActionFiles(),
		"rpc-private-key":             carapace.ActionFiles(),
		"save-cookies":                carapace.ActionFiles(),
		"save-session":                carapace.ActionFiles(),
		"server-stat-if":              carapace.ActionFiles(),
		"server-stat-of":              carapace.ActionFiles(),
		"stop-with-process":           ps.ActionProcessIds(),
		"stream-piece-selector":       carapace.ActionValues("default", "inorder", "random", "geom"),
		"torrent-file":                carapace.ActionFiles(".torrent"),
		"uri-selector":                carapace.ActionValues("inorder", "feedback", "adaptive"),
		"user-agent":                  http.ActionUserAgents(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
