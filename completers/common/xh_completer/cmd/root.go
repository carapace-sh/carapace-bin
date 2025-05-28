package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xh",
	Short: "Friendly and fast tool for sending HTTP requests",
	Long:  "https://github.com/ducaale/xh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("json", "j", true, "Serialize data items from the command line as a JSON object")
	rootCmd.Flags().BoolP("form", "f", false, "Serialize data items from the command line as form fields")
	rootCmd.Flags().Bool("multipart", false, "Like --form, but force a multipart/form-data request even without files")
	rootCmd.Flags().String("raw", "", "Pass raw request data without extra processing")
	rootCmd.Flags().String("pretty", "", "Controls output processing")
	rootCmd.Flags().String("format-options", "", "Set output formatting options")
	rootCmd.Flags().StringP("style", "s", "", "Output coloring style")
	rootCmd.Flags().String("response-charset", "", "Override the response encoding for terminal display purposes")
	rootCmd.Flags().String("response-mime", "", "Override the response mime type for coloring and formatting for the terminal")
	rootCmd.Flags().StringP("print", "p", "", "String specifying what the output should contain")
	rootCmd.Flags().BoolP("headers", "h", false, "Print only the response headers. Shortcut for --print=h")
	rootCmd.Flags().BoolP("body", "b", false, "Print only the response body. Shortcut for --print=b")
	rootCmd.Flags().BoolP("meta", "m", false, "Print only the response metadata. Shortcut for --print=m")
	rootCmd.Flags().BoolP("verbose", "v", false, "Print the whole request as well as the response")
	rootCmd.Flags().Bool("debug", false, "Print full error stack traces and debug log messages")
	rootCmd.Flags().Bool("all", false, "Show any intermediary requests/responses while following redirects with --follow")
	rootCmd.Flags().StringP("history-print", "P", "", "The same as --print but applies only to intermediary requests/responses")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not print to stdout or stderr")
	rootCmd.Flags().BoolP("stream", "S", false, "Always stream the response body")
	rootCmd.Flags().BoolP("compress", "x", false, "Content compressed (encoded) with Deflate algorithm")
	rootCmd.Flags().StringP("output", "o", "", "Save output to FILE instead of stdout")
	rootCmd.Flags().BoolP("download", "d", false, "Download the body to a file instead of printing it")

	// TODO: according to the description, --continue can only be used if --download and --output are used
	// how can we tell carapace to conditonally show flags based on existance of others?
	rootCmd.Flags().BoolP("continue", "c", false, "Resume an interrupted download. Requires --download and --output")
	rootCmd.Flags().String("session", "", "Create, or reuse and update a session")
	rootCmd.Flags().String("session-read-only", "", "Create or read a session without updating it form the request/response exchange")
	rootCmd.Flags().StringP("auth-type", "A", "", "Specify the auth mechanism")
	rootCmd.Flags().StringP("auth", "a", "", "Authenticate as USER with PASS (-A basic|digest) or with TOKEN (-A bearer)")
	rootCmd.Flags().Bool("ignore-netrc", false, "Do not use credentials from .netrc")
	rootCmd.Flags().Bool("offline", false, "Construct HTTP requests without sending them anywhere")
	rootCmd.Flags().Bool("check-status", true, "Exit with an error status code if the server replies with an error")
	rootCmd.Flags().BoolP("follow", "F", false, "Do follow redirects")

	// TODO: according to the description, --max-redirects can only be used if --follow is used
	// how can we tell carapace to conditonally show flags based on existance of others?
	rootCmd.Flags().Int("max-redirects", 0, "Number of redirects to follow. Only respected if --follow is used")
	rootCmd.Flags().Int("timeout", 0, "Connection timeout of the request")
	rootCmd.Flags().String("proxy", "", "Use a proxy for a protocol. For example: --proxy https:http://proxy.host:8080")
	rootCmd.Flags().String("verify", "", "If 'no', skip SSL verification. If a file path, use it as a CA bundle")
	rootCmd.Flags().String("cert", "", "Use a client side certificate for SSL")
	rootCmd.Flags().String("cert-key", "", "A private key file to use with --cert")
	rootCmd.Flags().String("ssl", "", "Force a particular TLS version")
	rootCmd.Flags().Bool("https", false, "Make HTTPS requests if not specified in the URL")
	rootCmd.Flags().String("http-version", "", "HTTP version to use")
	rootCmd.Flags().String("resolve", "", "Override DNS resolution for specific domain to a custom IP")
	rootCmd.Flags().String("interface", "", "Bind to a network interface or local IP address")
	rootCmd.Flags().StringP("ipv4", "4", "", "Resolve hostname to ipv4 addresses only")
	rootCmd.Flags().StringP("ipv6", "6", "", "Resolve hostname to ipv6 addresses only")
	rootCmd.Flags().BoolP("ignore-stdin", "I", false, "Do not attempt to read stdin")
	rootCmd.Flags().Bool("curl", false, "Print a translation to a curl command")
	rootCmd.Flags().Bool("curl-long", false, "Use the long versions of curl's flags")
	rootCmd.Flags().String("generate", "", "Generate shell completions or man pages")
	rootCmd.Flags().Bool("help", false, "Print help")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	rootCmd.MarkFlagsMutuallyExclusive("json", "form", "multipart")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"pretty":            carapace.ActionValues("all", "colors", "format", "none"),
		"style":             carapace.ActionValues("auto", "solarized", "monokai", "fruity"),
		"print":             carapace.ActionValues("h", "b", "m").UniqueList(""),
		"history-print":     carapace.ActionValues("h", "b", "m").UniqueList(""),
		"output":            carapace.ActionFiles(),
		"session":           carapace.ActionFiles(),
		"session-read-only": carapace.ActionFiles(),
		"auth-type":         carapace.ActionValues("basic", "bearer", "digest"),
		"verify": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			actionStatic := carapace.ActionValuesDescribed("no", "Skip SSL verification").Invoke(c)
			actionFiles := carapace.ActionFiles().Invoke(c)
			return actionStatic.Merge(actionFiles).ToA()
		}),
		"ssl":          carapace.ActionValues("auto", "tls1", "tls1.1", "tls1.2", "tls1.3"),
		"http-version": carapace.ActionValues("1.0", "1.1", "2", "2-prior-knowledge"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "TRACE", "CONNECT"),
	)
}
