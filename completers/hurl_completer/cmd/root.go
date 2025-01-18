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
	Short: "run and test HTTP requests with plain text",
	Long:  "https://hurl.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("aws-sigv4", "", "Use AWS V4 signature authentication in the transfer")
	rootCmd.Flags().String("cacert", "", "CA certificate to verify peer against")
	rootCmd.Flags().StringP("cert", "E", "", "Client certificate file and password")
	rootCmd.Flags().Bool("color", false, "Colorize output")
	rootCmd.Flags().Bool("compressed", false, "Request compressed response")
	rootCmd.Flags().String("connect-timeout", "", "Maximum time allowed for connection")
	rootCmd.Flags().Bool("connect-to", false, "For given HOST1:PORT1 pair, connect to HOST2:PORT2 instead")
	rootCmd.Flags().Bool("continue-on-error", false, "Continue executing requests even if an error occurs")
	rootCmd.Flags().StringP("cookie", "b", "", "Read cookies from FILE")
	rootCmd.Flags().StringP("cookie-jar", "c", "", "Write cookies to FILE after running the session")
	rootCmd.Flags().String("curl", "", "Export each request to a list of curl commands")
	rootCmd.Flags().String("delay", "", "Sets delay before each request")
	rootCmd.Flags().String("error-format", "", "Control the format of error messages")
	rootCmd.Flags().String("file-root", "", "Set root directory to import files")
	rootCmd.Flags().String("from-entry", "", "Execute Hurl file from ENTRY_NUMBER")
	rootCmd.Flags().String("glob", "", "Specify input files that match the given GLOB")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("http1.0", "0", false, "Use HTTP version 1.0")
	rootCmd.Flags().Bool("http1.1", false, "Use HTTP version 1.1")
	rootCmd.Flags().Bool("http2", false, "Use HTTP version 2")
	rootCmd.Flags().Bool("http3", false, "Use HTTP version 3")
	rootCmd.Flags().Bool("ignore-asserts", false, "Ignore asserts defined in the Hurl file")
	rootCmd.Flags().BoolP("include", "i", false, "Include the HTTP headers in the output")
	rootCmd.Flags().BoolP("insecure", "k", false, "Allow insecure SSL connections")
	rootCmd.Flags().Bool("interactive", false, "Turn on interactive mode")
	rootCmd.Flags().BoolP("ipv4", "4", false, "Use IPv4 addresses only when resolving host names")
	rootCmd.Flags().BoolP("ipv6", "6", false, "Use IPv6 addresses only when resolving host names")
	rootCmd.Flags().String("jobs", "", "Maximum number of parallel jobs")
	rootCmd.Flags().Bool("json", false, "Output each Hurl file result to JSON")
	rootCmd.Flags().String("key", "", "Private key file name")
	rootCmd.Flags().String("limit-rate", "", "Specify the maximum transfer rate in bytes/second")
	rootCmd.Flags().BoolP("location", "L", false, "Follow redirects")
	rootCmd.Flags().Bool("location-trusted", false, "Allow sending the name + password to redirected hosts")
	rootCmd.Flags().String("max-filesize", "", "Specify the maximum size in bytes of a file to download")
	rootCmd.Flags().String("max-redirs", "", "Maximum number of redirects allowed")
	rootCmd.Flags().StringP("max-time", "m", "", "Maximum time allowed for the transfer")
	rootCmd.Flags().BoolP("netrc", "n", false, "Must read .netrc for username and password")
	rootCmd.Flags().String("netrc-file", "", "Specify FILE for .netrc")
	rootCmd.Flags().Bool("netrc-optional", false, "Use either .netrc or the URL")
	rootCmd.Flags().Bool("no-color", false, "Do not colorize output")
	rootCmd.Flags().Bool("no-output", false, "Suppress output")
	rootCmd.Flags().String("noproxy", "", "List of hosts which do not use proxy")
	rootCmd.Flags().StringP("output", "o", "", "Write to FILE instead of stdout")
	rootCmd.Flags().Bool("parallel", false, "Run files in parallel")
	rootCmd.Flags().Bool("path-as-is", false, "Do not handle sequences of /../ or /./ in the given URL path")
	rootCmd.Flags().StringP("proxy", "x", "", "Use proxy on given PROTOCOL/HOST/PORT")
	rootCmd.Flags().String("repeat", "", "Repeat the input files sequence NUM times")
	rootCmd.Flags().String("report-html", "", "Generate HTML report to DIR")
	rootCmd.Flags().String("report-json", "", "Generate JSON report to DIR")
	rootCmd.Flags().String("report-junit", "", "Write a JUnit XML report to FILE")
	rootCmd.Flags().String("report-tap", "", "Write a TAP report to FILE")
	rootCmd.Flags().String("resolve", "", "Provide a custom address for a specific HOST and PORT pair")
	rootCmd.Flags().String("retry", "", "Maximum number of retries, 0 for no retries")
	rootCmd.Flags().String("retry-interval", "", "Interval in milliseconds before a retry")
	rootCmd.Flags().Bool("ssl-no-revoke", false, "Disable certificate revocation checks")
	rootCmd.Flags().Bool("test", false, "Activate test mode")
	rootCmd.Flags().String("to-entry", "", "Execute Hurl file to ENTRY_NUMBER")
	rootCmd.Flags().String("unix-socket", "", "Connect through this Unix domain socket")
	rootCmd.Flags().StringP("user", "u", "", "Add basic Authentication header to each request")
	rootCmd.Flags().StringP("user-agent", "A", "", "Specify the User-Agent string to send to the HTTP server")
	rootCmd.Flags().String("variable", "", "Define a variable")
	rootCmd.Flags().String("variables-file", "", "Define a properties file in which you define your variables")
	rootCmd.Flags().BoolP("verbose", "v", false, "Turn on verbose")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().Bool("very-verbose", false, "Turn on verbose output")

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
		"cookie":       carapace.ActionFiles(),
		"cookie-jar":   carapace.ActionFiles(),
		"curl":         carapace.ActionFiles(),
		"error-format": carapace.ActionValues("short", "long"),
		"file-root":    carapace.ActionDirectories(),
		"key":          carapace.ActionFiles(),
		"netrc-file":   carapace.ActionFiles(),
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
		"unix-socket":    carapace.ActionFiles(),
		"user-agent":     http.ActionUserAgents(),
		"variables-file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
