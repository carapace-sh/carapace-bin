package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "minikube",
	Short: "minikube quickly sets up a local Kubernetes cluster",
	Long:  "https://minikube.sigs.k8s.io/docs/",
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

	rootCmd.AddGroup(
		&cobra.Group{ID: "basic", Title: "Basic Commands"},
		&cobra.Group{ID: "images", Title: "Images Commands"},
		&cobra.Group{ID: "configuration", Title: "Configuration and Management Commands"},
		&cobra.Group{ID: "networking", Title: "Networking and Connectivity Commands"},
		&cobra.Group{ID: "advanced", Title: "Advanced Commands"},
		&cobra.Group{ID: "troubleshooting", Title: "Troubleshooting Commands"},
	)

	rootCmd.PersistentFlags().Bool("add_dir_header", false, "If true, adds the file directory to the header of the log messages")
	rootCmd.PersistentFlags().Bool("alsologtostderr", false, "log to standard error as well as files (no effect when -logtostderr=true)")
	rootCmd.PersistentFlags().StringP("bootstrapper", "b", "", "The name of the cluster bootstrapper that will set up the Kubernetes cluster.")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "")
	rootCmd.PersistentFlags().String("log_backtrace_at", "", "when logging hits line file:N, emit a stack trace")
	rootCmd.PersistentFlags().String("log_dir", "", "If non-empty, write log files in this directory (no effect when -logtostderr=true)")
	rootCmd.PersistentFlags().String("log_file", "", "If non-empty, use this log file (no effect when -logtostderr=true)")
	rootCmd.PersistentFlags().String("log_file_max_size", "", "Defines the maximum size a log file can grow to (no effect when -logtostderr=true). Unit is megabytes. If the value is 0, the maximum file size is unlimited.")
	rootCmd.PersistentFlags().Bool("logtostderr", false, "log to standard error instead of files")
	rootCmd.PersistentFlags().Bool("one_output", false, "If true, only write logs to their native severity level (vs also writing to each lower severity level; no effect when -logtostderr=true)")
	rootCmd.PersistentFlags().StringP("profile", "p", "", "The name of the minikube VM being used. This can be set to allow having multiple instances of minikube independently.")
	rootCmd.PersistentFlags().Bool("rootless", false, "Force to use rootless driver (docker and podman driver only)")
	rootCmd.PersistentFlags().Bool("skip-audit", false, "Skip recording the current command in the audit logs.")
	rootCmd.PersistentFlags().Bool("skip_headers", false, "If true, avoid header prefixes in the log messages")
	rootCmd.PersistentFlags().Bool("skip_log_headers", false, "If true, avoid headers when opening log files (no effect when -logtostderr=true)")
	rootCmd.PersistentFlags().String("stderrthreshold", "", "logs at or above this threshold go to stderr when writing to files and stderr (no effect when -logtostderr=true or -alsologtostderr=true)")
	rootCmd.PersistentFlags().String("user", "", "Specifies the user executing the operation. Useful for auditing operations executed by 3rd party tools. Defaults to the operating system username.")
	rootCmd.PersistentFlags().StringP("v", "v", "", "number for the log level verbosity")
	rootCmd.PersistentFlags().String("vmodule", "", "comma-separated list of pattern=N settings for file-filtered logging")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log_dir":  carapace.ActionDirectories(),
		"log_file": carapace.ActionFiles(),
		"user":     os.ActionUsers(),
	})
}
