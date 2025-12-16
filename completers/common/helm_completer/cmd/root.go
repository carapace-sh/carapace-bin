package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helm",
	Short: "The Helm package manager for Kubernetes.",
	Long:  "https://helm.sh/",
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
		&cobra.Group{ID: "main", Title: "Main Commands"},
		&cobra.Group{ID: "plugin", Title: "Plugin Commands"},
	)

	rootCmd.PersistentFlags().Bool("add-dir-header", false, "If true, adds the file directory to the header of the log messages")
	rootCmd.PersistentFlags().Bool("alsologtostderr", false, "log to standard error as well as files")
	rootCmd.PersistentFlags().Bool("debug", false, "enable verbose output")
	rootCmd.Flags().BoolP("help", "h", false, "help for helm")
	rootCmd.PersistentFlags().String("kube-apiserver", "", "the address and the port for the Kubernetes API server")
	rootCmd.PersistentFlags().StringArray("kube-as-group", nil, "group to impersonate for the operation, this flag can be repeated to specify multiple groups.")
	rootCmd.PersistentFlags().String("kube-as-user", "", "username to impersonate for the operation")
	rootCmd.PersistentFlags().String("kube-ca-file", "", "the certificate authority file for the Kubernetes API server connection")
	rootCmd.PersistentFlags().String("kube-context", "", "name of the kubeconfig context to use")
	rootCmd.PersistentFlags().String("kube-token", "", "bearer token used for authentication")
	rootCmd.PersistentFlags().String("kubeconfig", "", "path to the kubeconfig file")
	rootCmd.PersistentFlags().String("log-backtrace-at", ":0", "when logging hits line file:N, emit a stack trace")
	rootCmd.PersistentFlags().String("log-dir", "", "If non-empty, write log files in this directory")
	rootCmd.PersistentFlags().String("log-file", "", "If non-empty, use this log file")
	rootCmd.PersistentFlags().Uint64("log-file-max-size", 1800, "Defines the maximum size a log file can grow to. Unit is megabytes. If the value is 0, the maximum file size is unlimited.")
	rootCmd.PersistentFlags().Bool("logtostderr", true, "log to standard error instead of files")
	rootCmd.PersistentFlags().StringP("namespace", "n", "", "namespace scope for this request")
	rootCmd.PersistentFlags().Bool("one-output", false, "If true, only write logs to their native severity level (vs also writing to each lower severity level)")
	rootCmd.PersistentFlags().String("registry-config", "/home/user/.config/helm/registry.json", "path to the registry config file")
	rootCmd.PersistentFlags().String("repository-cache", "/home/user/.cache/helm/repository", "path to the file containing cached repository indexes")
	rootCmd.PersistentFlags().String("repository-config", "/home/user/.config/helm/repositories.yaml", "path to the file containing repository names and URLs")
	rootCmd.PersistentFlags().Bool("skip-headers", false, "If true, avoid header prefixes in the log messages")
	rootCmd.PersistentFlags().Bool("skip-log-headers", false, "If true, avoid headers when opening log files")
	rootCmd.PersistentFlags().Int("stderrthreshold", 2, "logs at or above this threshold go to stderr")
	rootCmd.PersistentFlags().IntP("v", "v", 0, "number for the log level verbosity")
	rootCmd.PersistentFlags().StringSlice("vmodule", nil, "comma-separated list of pattern=N settings for file-filtered logging")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"kube-as-group": os.ActionGroups(),
		"kube-as-user":  os.ActionUsers(),
		"kube-ca-file":  carapace.ActionFiles(),
		"kube-context":  kubectl.ActionContexts(),
		"kubeconfig":    carapace.ActionFiles(),
		"log-dir":       carapace.ActionDirectories(),
		"log-file":      carapace.ActionFiles(),
		"namespace": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{Types: "namespaces", Context: rootCmd.Flag("kube-context").Value.String()})
		}),
		"registry-config":  carapace.ActionFiles(),
		"repository-cache": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if _, _, err := cmd.Find(args); len(args) > 1 && err == nil {
			return // core command - skip plugin commands
		}
		addPluginCommands()
	})
}

func addPluginCommands() {
	if output, err := (carapace.Context{}).Command("helm", "plugin", "list").Output(); err == nil {
		lines := strings.Split(string(output), "\n")

		if len(lines) < 2 {
			return
		}

		for _, line := range lines[1 : len(lines)-1] {
			if splitted := strings.SplitN(line, "\t", 3); len(splitted) == 3 {
				pluginCmd := &cobra.Command{
					Use:                splitted[0],
					Short:              splitted[2],
					Run:                func(cmd *cobra.Command, args []string) {},
					GroupID:            "plugin",
					DisableFlagParsing: true,
				}

				carapace.Gen(pluginCmd).PositionalAnyCompletion(
					bridge.ActionCarapaceBin("helm-" + splitted[0]),
				)

				rootCmd.AddCommand(pluginCmd)
			}
		}
	}
}
