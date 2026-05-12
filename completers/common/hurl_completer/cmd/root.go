package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hurl",
	Short: "Hurl, run and test HTTP requests with plain text",
	Long:  "https://hurl.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("aws-sigv4", "", "Use AWS V4 signature authentication in the transfer")
	rootCmd.Flags().String("cacert", "", "CA certificate to verify peer against (PEM format)")
	rootCmd.Flags().StringP("cert", "E", "", "Client certificate file and password")
	rootCmd.Flags().Bool("color", false, "Colorize output")
	rootCmd.Flags().Bool("compressed", false, "Request compressed response (using deflate or gzip)")
	rootCmd.Flags().String("connect-timeout", "", "Maximum time allowed for connection [default: 300]")
	rootCmd.Flags().StringSlice("connect-to", nil, "For a request to the given HOST1:PORT1 pair, connect to HOST2:PORT2 instead")
	rootCmd.Flags().Bool("continue-on-error", false, "Continue executing requests even if an error occurs")
	rootCmd.Flags().StringP("cookie", "b", "", "Read cookies from FILE")
	rootCmd.Flags().StringP("cookie-jar", "c", "", "Write cookies to FILE after running the session")
	rootCmd.Flags().String("curl", "", "Export each request to a list of curl commands")
	rootCmd.Flags().String("delay", "", "Sets delay before each request (aka sleep) [default: 0]")
	rootCmd.Flags().Bool("digest", false, "Tell Hurl to use HTTP Digest authentication")
	rootCmd.Flags().String("error-format", "", "Control the format of error messages [default: short]")
	rootCmd.Flags().String("file-root", "", "Set root directory to import files [default: input file directory]")
	rootCmd.Flags().String("from-entry", "", "Execute Hurl file from ENTRY_NUMBER (starting at 1)")
	rootCmd.Flags().StringSlice("glob", nil, "Specify input files that match the given GLOB. Multiple glob flags may be used")
	rootCmd.Flags().StringSliceP("header", "H", nil, "Pass custom header(s) to server")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("http1.0", "0", false, "Tell Hurl to use HTTP version 1.0")
	rootCmd.Flags().Bool("http1.1", false, "Tell Hurl to use HTTP version 1.1")
	rootCmd.Flags().Bool("http2", false, "Tell Hurl to use HTTP version 2")
	rootCmd.Flags().Bool("http3", false, "Tell Hurl to use HTTP version 3")
	rootCmd.Flags().BoolP("include", "i", false, "Include the HTTP headers in the output")
	rootCmd.Flags().BoolP("insecure", "k", false, "Allow insecure SSL connections")
	rootCmd.Flags().BoolP("ipv4", "4", false, "Tell Hurl to use IPv4 addresses only when resolving host names, and not for example try IPv6")
	rootCmd.Flags().BoolP("ipv6", "6", false, "Tell Hurl to use IPv6 addresses only when resolving host names, and not for example try IPv4")
	rootCmd.Flags().String("jobs", "", "Maximum number of parallel jobs, 1 to disable parallel execution")
	rootCmd.Flags().Bool("json", false, "Output each Hurl file result to JSON")
	rootCmd.Flags().String("key", "", "Private key file name")
	rootCmd.Flags().String("limit-rate", "", "Specify the maximum transfer rate in bytes/second, for both downloads and uploads")
	rootCmd.Flags().BoolP("location", "L", false, "Follow redirects")
	rootCmd.Flags().Bool("location-trusted", false, "Follow redirects but allows sending the name + password to all hosts that the site may redirect to")
	rootCmd.Flags().String("max-filesize", "", "Specify the maximum size in bytes of a file to download")
	rootCmd.Flags().String("max-redirs", "", "Maximum number of redirects allowed, -1 for unlimited redirects [default: 50]")
	rootCmd.Flags().StringP("max-time", "m", "", "Maximum time allowed for the transfer [default: 300]")
	rootCmd.Flags().Bool("negotiate", false, "Tell Hurl to use Negotiate (SPNEGO) authentication")
	rootCmd.Flags().BoolP("netrc", "n", false, "Must read .netrc for username and password")
	rootCmd.Flags().String("netrc-file", "", "Specify FILE for .netrc")
	rootCmd.Flags().Bool("netrc-optional", false, "Use either .netrc or the URL")
	rootCmd.Flags().Bool("no-assert", false, "Ignore asserts defined in the Hurl file")
	rootCmd.Flags().Bool("no-color", false, "Do not colorize output")
	rootCmd.Flags().Bool("no-cookie-store", false, "Do not use cookie store between requests")
	rootCmd.Flags().Bool("no-output", false, "Suppress output. By default, Hurl outputs the body of the last response")
	rootCmd.Flags().Bool("no-pretty", false, "Do not prettify response output")
	rootCmd.Flags().String("no-proxy", "", "List of hosts which do not use proxy")
	rootCmd.Flags().String("noproxy", "", "List of hosts which do not use proxy")
	rootCmd.Flags().Bool("ntlm", false, "Tell Hurl to use NTLM authentication")
	rootCmd.Flags().StringP("output", "o", "", "Write to FILE instead of stdout")
	rootCmd.Flags().Bool("parallel", false, "Run files in parallel (default in test mode)")
	rootCmd.Flags().Bool("path-as-is", false, "Tell Hurl to not handle sequences of /../ or /./ in the given URL path")
	rootCmd.Flags().String("pinnedpubkey", "", "Public key to verify peer against")
	rootCmd.Flags().Bool("pretty", false, "Prettify JSON response output")
	rootCmd.Flags().Bool("progress-bar", false, "Display a progress bar in test mode")
	rootCmd.Flags().StringP("proxy", "x", "", "Use proxy on given PROTOCOL/HOST/PORT")
	rootCmd.Flags().String("repeat", "", "Repeat the input files sequence NUM times, -1 for infinite loop")
	rootCmd.Flags().String("report-html", "", "Generate HTML report to DIR")
	rootCmd.Flags().String("report-json", "", "Generate JSON report to DIR")
	rootCmd.Flags().String("report-junit", "", "Write a JUnit XML report to FILE")
	rootCmd.Flags().String("report-tap", "", "Write a TAP report to FILE")
	rootCmd.Flags().StringSlice("resolve", nil, "Provide a custom address for a specific HOST and PORT pair")
	rootCmd.Flags().String("retry", "", "Maximum number of retries, 0 for no retries, -1 for unlimited retries")
	rootCmd.Flags().String("retry-interval", "", "Interval in milliseconds before a retry [default: 1000]")
	rootCmd.Flags().StringSlice("secret", nil, "Define a variable which value is secret")
	rootCmd.Flags().StringSlice("secrets-file", nil, "Define a secrets file in which you define your secrets")
	rootCmd.Flags().Bool("ssl-no-revoke", false, "(Windows) Tell Hurl to disable certificate revocation checks")
	rootCmd.Flags().Bool("test", false, "Activate test mode (use parallel execution)")
	rootCmd.Flags().String("to-entry", "", "Execute Hurl file to ENTRY_NUMBER (starting at 1)")
	rootCmd.Flags().String("unix-socket", "", "(HTTP) Connect through this Unix domain socket, instead of using the network")
	rootCmd.Flags().StringP("user", "u", "", "Add basic Authentication header to each request")
	rootCmd.Flags().StringP("user-agent", "A", "", "Specify the User-Agent string to send to the HTTP server")
	rootCmd.Flags().StringSlice("variable", nil, "Define a variable")
	rootCmd.Flags().StringSlice("variables-file", nil, "Define a properties file in which you define your variables")
	rootCmd.Flags().BoolP("verbose", "v", false, "Turn on verbose (alias to --verbosity verbose)")
	rootCmd.Flags().String("verbosity", "", "Set verbosity level for debug log")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().Bool("very-verbose", false, "Turn on very verbose output, including HTTP response and libcurl logs (alias to --verbosity debug)")
	rootCmd.Flag("noproxy").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"aws-sigv4": carapace.ActionValues(), // TODO
		"cacert":    carapace.ActionFiles(),
		"cert": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()

			}
		}),
		"connect-to": carapace.ActionMultiPartsN(":", 4, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts().Suffix(":")
			case 1:
				return net.ActionPorts().Suffix(":")
			case 2:
				return net.ActionHosts().Suffix(":")
			default:
				return net.ActionPorts()
			}
		}),
		"cookie":       carapace.ActionFiles(),
		"cookie-jar":   carapace.ActionFiles(),
		"curl":         carapace.ActionFiles(),
		"error-format": carapace.ActionValues("short", "long"),
		"file-root":    carapace.ActionDirectories(),
		"header":       http.ActionRequestHeaders(),
		"key":          carapace.ActionFiles(),
		"netrc-file":   carapace.ActionFiles(),
		"no-proxy":     net.ActionHosts().UniqueList(","),
		"noproxy":      net.ActionHosts(),
		"output":       carapace.ActionFiles(),
		"proxy": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !strings.Contains(c.Value, ":") {
				return carapace.Batch(
					net.ActionProtocols().Suffix("://"),
					net.ActionHosts(),
				).ToA().NoSpace()
			}

			prefix := ""
			if index := strings.Index(c.Value, "://"); index > -1 {
				prefix = c.Value[:index+3]
				c.Value = strings.TrimPrefix(c.Value, prefix)
			}

			return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return net.ActionHosts().NoSpace()
				default:
					return net.ActionPorts()
				}
			}).Invoke(c).Prefix(prefix).ToA()
		}),
		"report-html":  carapace.ActionDirectories(),
		"report-json":  carapace.ActionDirectories(),
		"report-junit": carapace.ActionFiles(),
		"report-tap":   carapace.ActionFiles(),
		"resolve": carapace.ActionMultiPartsN(":", 3, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts().Suffix(":")
			case 1:
				return net.ActionPorts().Suffix(":")
			default:
				return carapace.ActionValues()
			}
		}),
		"secrets-file":   carapace.ActionFiles(),
		"unix-socket":    carapace.ActionFiles(),
		"user-agent":     http.ActionUserAgents(),
		"variables-file": carapace.ActionFiles(),
		"verbosity":      carapace.ActionValues("brief", "verbose", "debug"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
