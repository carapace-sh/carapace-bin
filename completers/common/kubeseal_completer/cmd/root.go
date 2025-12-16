package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubeseal",
	Short: "A Kubernetes controller and tool for one-way encrypted Secrets",
	Long:  "https://github.com/bitnami-labs/sealed-secrets/",
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

	// Logging flags
	rootCmd.Flags().Bool("add_dir_header", false, "If true, adds the file directory to the header of the log messages")
	rootCmd.Flags().Bool("alsologtostderr", false, "log to standard error as well as files (no effect when -logtostderr=true)")
	rootCmd.Flags().String("log_backtrace_at", ":0", "when logging hits line file:N, emit a stack trace")
	rootCmd.Flags().String("log_dir", "", "If non-empty, write log files in this directory (no effect when -logtostderr=true)")
	rootCmd.Flags().String("log_file", "", "If non-empty, use this log file (no effect when -logtostderr=true)")
	rootCmd.Flags().Uint("log_file_max_size", 1800, "Defines the maximum size a log file can grow to (no effect when -logtostderr=true). Unit is megabytes. If the value is 0, the maximum file size is unlimited.")
	rootCmd.Flags().Bool("logtostderr", true, "log to standard error instead of files")
	rootCmd.Flags().Bool("one_output", false, "If true, only write logs to their native severity level (vs also writing to each lower severity level; no effect when -logtostderr=true)")
	rootCmd.Flags().Bool("skip_headers", false, "If true, avoid header prefixes in the log messages")
	rootCmd.Flags().Bool("skip_log_headers", false, "If true, avoid headers when opening log files (no effect when -logtostderr=true)")
	rootCmd.Flags().Int("stderrthreshold", 2, "logs at or above this threshold go to stderr when writing to files and stderr (no effect when -logtostderr=true or -alsologtostderr=true)")
	rootCmd.Flags().StringP("v", "v", "", "number for the log level verbosity")
	rootCmd.Flags().String("vmodule", "", "comma-separated list of pattern=N settings for file-filtered logging")

	// Kubernetes client flags
	rootCmd.Flags().String("as", "", "Username to impersonate for the operation")
	rootCmd.Flags().StringArray("as-group", nil, "Group to impersonate for the operation, this flag can be repeated to specify multiple groups.")
	rootCmd.Flags().String("as-uid", "", "UID to impersonate for the operation")
	rootCmd.Flags().String("certificate-authority", "", "Path to a cert file for the certificate authority")
	rootCmd.Flags().String("client-certificate", "", "Path to a client certificate file for TLS")
	rootCmd.Flags().String("client-key", "", "Path to a client key file for TLS")
	rootCmd.Flags().String("cluster", "", "The name of the kubeconfig cluster to use")
	rootCmd.Flags().String("context", "", "The name of the kubeconfig context to use")
	rootCmd.Flags().Bool("disable-compression", false, "If true, opt-out of response compression for all requests to the server")
	rootCmd.Flags().Bool("insecure-skip-tls-verify", false, "If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure")
	rootCmd.Flags().String("kubeconfig", "", "Path to a kube config. Only required if out-of-cluster")
	rootCmd.Flags().StringP("namespace", "n", "", "If present, the namespace scope for this CLI request")
	rootCmd.Flags().String("password", "", "Password for basic authentication to the API server")
	rootCmd.Flags().String("proxy-url", "", "If provided, this URL will be used to connect via proxy")
	rootCmd.Flags().String("request-timeout", "0", "The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests.")
	rootCmd.Flags().String("server", "", "The address and port of the Kubernetes API server")
	rootCmd.Flags().String("tls-server-name", "", "If provided, this name will be used to validate server certificate. If this is not provided, hostname used to contact the server is used.")
	rootCmd.Flags().String("token", "", "Bearer token for authentication to the API server")
	rootCmd.Flags().String("user", "", "The name of the kubeconfig user to use")
	rootCmd.Flags().String("username", "", "Username for basic authentication to the API server")

	// Kubeseal specific flags
	rootCmd.Flags().Bool("allow-empty-data", false, "Allow empty data in the secret object")
	rootCmd.Flags().String("cert", "", "Certificate / public key file/URL to use for encryption. Overrides --controller-*")
	rootCmd.Flags().String("controller-name", "sealed-secrets-controller", "Name of sealed-secrets controller.")
	rootCmd.Flags().String("controller-namespace", "kube-system", "Namespace of sealed-secrets controller.")
	rootCmd.Flags().Bool("fetch-cert", false, "Write certificate to stdout. Useful for later use with --cert")
	rootCmd.Flags().StringP("format", "o", "json", "Output format for sealed secret. Either json or yaml")
	rootCmd.Flags().StringSlice("from-file", nil, "(only with --raw) Secret items can be sourced from files. Pro-tip: you can use /dev/stdin to read pipe input. This flag tries to follow the same syntax as in kubectl")
	rootCmd.Flags().Bool("help", false, "Print this help message")
	rootCmd.Flags().String("merge-into", "", "Merge items from secret into an existing sealed secret file, updating the file in-place instead of writing to stdout.")
	rootCmd.Flags().String("name", "", "Name of the sealed secret (required with --raw and default (strict) scope)")
	rootCmd.Flags().Bool("raw", false, "Encrypt a raw value passed via the --from-* flags instead of the whole secret object")
	rootCmd.Flags().Bool("re-encrypt", false, "Re-encrypt the given sealed secret to use the latest cluster key.")
	rootCmd.Flags().StringSlice("recovery-private-key", nil, "Private key filename used by the --recovery-unseal command. Multiple files accepted either via comma separated list or by repetition of the flag. Either PEM encoded private keys or a backup of a json/yaml encoded k8s sealed-secret controller secret (and v1.List) are accepted.")
	rootCmd.Flags().Bool("recovery-unseal", false, "Decrypt a sealed secrets file obtained from stdin, using the private key passed with --recovery-private-key. Intended to be used in disaster recovery mode.")
	rootCmd.Flags().String("scope", "strict", "Set the scope of the sealed secret: strict, namespace-wide, cluster-wide (defaults to strict). Mandatory for --raw, otherwise the 'sealedsecrets.bitnami.com/cluster-wide' and 'sealedsecrets.bitnami.com/namespace-wide' annotations on the input secret can be used to select the scope.")
	rootCmd.Flags().StringP("sealed-secret-file", "w", "", "Sealed-secret (output) file")
	rootCmd.Flags().StringP("secret-file", "f", "", "Secret (input) file")
	rootCmd.Flags().Bool("validate", false, "Validate that the sealed secret can be decrypted")
	rootCmd.Flags().Bool("version", false, "Print version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cert":                  carapace.ActionFiles(),
		"certificate-authority": carapace.ActionFiles(),
		"client-certificate":    carapace.ActionFiles(),
		"client-key":            carapace.ActionFiles(),
		"context":               kubectl.ActionContexts(),
		"format":                carapace.ActionValues("json", "yaml"),
		"from-file":             carapace.ActionFiles(),
		"kubeconfig":            carapace.ActionFiles(),
		"log_dir":               carapace.ActionDirectories(),
		"log_file":              carapace.ActionFiles(),
		"merge-into":            carapace.ActionFiles(),
		"namespace": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Context: rootCmd.Flag("context").Value.String(),
				Types:   "namespaces",
			})
		}),
		"recovery-private-key": carapace.ActionFiles(),
		"scope":                carapace.ActionValues("strict", "namespace-wide", "cluster-wide"),
		"sealed-secret-file":   carapace.ActionFiles(),
		"secret-file":          carapace.ActionFiles(),
	})
}
