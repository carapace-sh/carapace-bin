package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("add-dir-header", false, "adds the file directory to the header of the log messages")
	rootCmd.Flags().Bool("alsologtostderr", false, "log to standard error as well as files")
	rootCmd.Flags().String("as", "", "Username to impersonate for the operation")
	rootCmd.Flags().String("as-group", "", "Group to impersonate for the operation")
	rootCmd.Flags().String("cache-dir", "", "Default cache directory")
	rootCmd.Flags().String("certificate-authority", "", "Path to a cert file for the certificate authority")
	rootCmd.Flags().String("client-certificate", "", "Path to a client certificate file for TLS")
	rootCmd.Flags().String("client-key", "", "Path to a client key file for TLS")
	rootCmd.Flags().String("cluster", "", "The name of the kubeconfig cluster to use")
	rootCmd.Flags().String("context", "", "The name of the kubeconfig context to use")
	rootCmd.Flags().Bool("insecure-skip-tls-verify", false, "don't check server's certificate for validity")
	rootCmd.Flags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	rootCmd.Flags().String("log-backtrace-at", "", "when logging hits line file:N, emit a stack trace")
	rootCmd.Flags().String("log-dir", "", "write log files in this directory")
	rootCmd.Flags().String("log-file", "", "use this log file")
	rootCmd.Flags().String("log-file-max-size", "", "maximum size a log file can grow to")
	rootCmd.Flags().String("log-flush-frequency", "", "Maximum number of seconds between log flushes")
	rootCmd.Flags().Bool("logtostderr", false, "log to standard error instead of files")
	rootCmd.Flags().Bool("match-server-version", false, "Require server version to match client version")
	rootCmd.Flags().String("namespace", "", "the namespace scope for this CLI request")
	rootCmd.Flags().String("password", "", "Password for basic authentication to the API server")
	rootCmd.Flags().String("profile", "", "Name of profile to capture")
	rootCmd.Flags().String("profile-output", "", "Name of the file to write the profile to")
	rootCmd.Flags().String("request-timeout", "", "The length of time to wait before giving up on a single server request")
	rootCmd.Flags().String("server", "", "The address and port of the Kubernetes API server")
	rootCmd.Flags().Bool("skip-headers", false, "avoid header prefixes in the log messages")
	rootCmd.Flags().Bool("skip-log-headers", false, "avoid headers when opening log files")
	rootCmd.Flags().String("stderrthreshold", "", "logs at or above this threshold go to stderr")
	rootCmd.Flags().String("tls-server-name", "", "Server name to use for server certificate validation")
	rootCmd.Flags().String("token", "", "Bearer token for authentication to the API server")
	rootCmd.Flags().String("user", "", "The name of the kubeconfig user to use")
	rootCmd.Flags().String("username", "", "Username for basic authentication to the API server")
	rootCmd.Flags().String("vmodule", "", "comma-separated list of pattern=N settings for file-filtered logging")
	rootCmd.Flags().Bool("warnings-as-errors", false, "Treat warnings received from the server as errors")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir":             carapace.ActionDirectories(),
		"certificate-authority": carapace.ActionFiles(""),
		"client-certificate":    carapace.ActionFiles(""),
		"client-key":            carapace.ActionFiles(""),
		"kubeconfig":            carapace.ActionFiles(""),
		"log-dir":               carapace.ActionDirectories(),
		"log-file":              carapace.ActionFiles(""),
		"profile-output":        carapace.ActionFiles(""),
		// TODO add completions for kubeconfig based flags
	})
}
