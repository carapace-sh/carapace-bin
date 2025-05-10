package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix",
	Short: "a tool for reproducible and declarative configuration management",
	Long:  "https://github.com/NixOS/nix",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "main", Title: "Main commands"},
		&cobra.Group{ID: "infrequently used", Title: "Infrequently used commands"},
		&cobra.Group{ID: "utility", Title: "Utility commands"},
		&cobra.Group{ID: "troubleshooting", Title: "Troubleshooting commands"},
	)

	rootCmd.Flags().Bool("accept-flake-config", false, "Enable the accept-flake-config setting")
	rootCmd.Flags().String("access-tokens", "", "Set the access-tokens setting")
	rootCmd.Flags().Bool("allow-dirty", false, "Enable the allow-dirty setting")
	rootCmd.Flags().Bool("allow-import-from-derivation", false, "Enable the allow-import-from-derivation setting")
	rootCmd.Flags().Bool("allow-new-privileges", false, "Enable the allow-new-privileges setting")
	rootCmd.Flags().Bool("allow-symlinked-store", false, "Enable the allow-symlinked-store setting")
	rootCmd.Flags().Bool("allow-unsafe-native-code-during-evaluation", false, "Enable the allow-unsafe-native-code-during-evaluation setting")
	rootCmd.Flags().String("allowed-impure-host-deps", "", "Set the allowed-impure-host-deps setting")
	rootCmd.Flags().String("allowed-uris", "", "Set the allowed-uris setting")
	rootCmd.Flags().String("allowed-users", "", "Set the allowed-users setting")
	rootCmd.Flags().Bool("auto-optimise-store", false, "Enable the auto-optimise-store setting")
	rootCmd.Flags().String("bash-prompt", "", "Set the bash-prompt setting")
	rootCmd.Flags().String("bash-prompt-prefix", "", "Set the bash-prompt-prefix setting")
	rootCmd.Flags().String("bash-prompt-suffix", "", "Set the bash-prompt-suffix setting")
	rootCmd.Flags().String("build-hook", "", "Set the build-hook setting")
	rootCmd.Flags().String("build-poll-interval", "", "Set the build-poll-interval setting")
	rootCmd.Flags().String("build-users-group", "", "Set the build-users-group setting")
	rootCmd.Flags().StringArray("builders", nil, "Set the builders setting")
	rootCmd.Flags().Bool("builders-use-substitutes", false, "Enable the builders-use-substitutes setting")
	rootCmd.Flags().String("commit-lockfile-summary", "", "Set the commit-lockfile-summary setting")
	rootCmd.Flags().Bool("compress-build-log", false, "Enable the compress-build-log setting")
	rootCmd.Flags().String("connect-timeout", "", "Set the connect-timeout setting")
	rootCmd.Flags().String("cores", "", "Set the cores setting")
	rootCmd.Flags().Bool("debug", false, "Set the logging verbosity level to ‘debug’")
	rootCmd.Flags().String("diff-hook", "", "Set the diff-hook setting")
	rootCmd.Flags().String("download-attempts", "", "Set the download-attempts setting")
	rootCmd.Flags().String("download-speed", "", "Set the download-speed setting")
	rootCmd.Flags().Bool("enforce-determinism", false, "Enable the enforce-determinism setting")
	rootCmd.Flags().Bool("eval-cache", false, "Enable the eval-cache setting")
	rootCmd.Flags().String("experimental-features", "", "Set the experimental-features setting")
	rootCmd.Flags().String("extra-access-tokens", "", "Append to the access-tokens setting")
	rootCmd.Flags().String("extra-allowed-impure-host-deps", "", "Append to the allowed-impure-host-deps setting")
	rootCmd.Flags().String("extra-allowed-uris", "", "Append to the allowed-uris setting")
	rootCmd.Flags().String("extra-allowed-users", "", "Append to the allowed-users setting")
	rootCmd.Flags().String("extra-experimental-features", "", "Append to the experimental-features setting")
	rootCmd.Flags().String("extra-extra-platforms", "", "Append to the extra-platforms setting")
	rootCmd.Flags().String("extra-hashed-mirrors", "", "Append to the hashed-mirrors setting")
	rootCmd.Flags().String("extra-ignored-acls", "", "Append to the ignored-acls setting")
	rootCmd.Flags().String("extra-nix-path", "", "Append to the nix-path setting")
	rootCmd.Flags().String("extra-platforms", "", "Set the extra-platforms setting")
	rootCmd.Flags().String("extra-plugin-files", "", "Append to the plugin-files setting")
	rootCmd.Flags().String("extra-sandbox-paths", "", "Append to the sandbox-paths setting")
	rootCmd.Flags().String("extra-secret-key-files", "", "Append to the secret-key-files setting")
	rootCmd.Flags().String("extra-substituters", "", "Append to the substituters setting")
	rootCmd.Flags().String("extra-system-features", "", "Append to the system-features setting")
	rootCmd.Flags().String("extra-trusted-public-keys", "", "Append to the trusted-public-keys setting")
	rootCmd.Flags().String("extra-trusted-substituters", "", "Append to the trusted-substituters setting")
	rootCmd.Flags().String("extra-trusted-users", "", "Append to the trusted-users setting")
	rootCmd.Flags().Bool("fallback", false, "Enable the fallback setting")
	rootCmd.Flags().Bool("filter-syscalls", false, "Enable the filter-syscalls setting")
	rootCmd.Flags().String("flake-registry", "", "Set the flake-registry setting")
	rootCmd.Flags().Bool("fsync-metadata", false, "Enable the fsync-metadata setting")
	rootCmd.Flags().String("gc-reserved-space", "", "Set the gc-reserved-space setting")
	rootCmd.Flags().String("hashed-mirrors", "", "Set the hashed-mirrors setting")
	rootCmd.Flags().String("http-connections", "", "Set the http-connections setting")
	rootCmd.Flags().Bool("http2", false, "Enable the http2 setting")
	rootCmd.Flags().Bool("ignore-try", false, "Enable the ignore-try setting")
	rootCmd.Flags().String("ignored-acls", "", "Set the ignored-acls setting")
	rootCmd.Flags().Bool("impersonate-linux-26", false, "Enable the impersonate-linux-26 setting")
	rootCmd.Flags().Bool("keep-build-log", false, "Enable the keep-build-log setting")
	rootCmd.Flags().Bool("keep-derivations", false, "Enable the keep-derivations setting")
	rootCmd.Flags().Bool("keep-env-derivations", false, "Enable the keep-env-derivations setting")
	rootCmd.Flags().Bool("keep-failed", false, "Enable the keep-failed setting")
	rootCmd.Flags().Bool("keep-going", false, "Enable the keep-going setting")
	rootCmd.Flags().Bool("keep-outputs", false, "Enable the keep-outputs setting")
	rootCmd.Flags().String("log-format", "", "Set the format of log output")
	rootCmd.Flags().String("log-lines", "", "Set the log-lines setting")
	rootCmd.Flags().String("max-build-log-size", "", "Set the max-build-log-size setting")
	rootCmd.Flags().String("max-free", "", "Set the max-free setting")
	rootCmd.Flags().String("max-jobs", "", "Set the max-jobs setting")
	rootCmd.Flags().String("max-silent-time", "", "Set the max-silent-time setting")
	rootCmd.Flags().String("min-free", "", "Set the min-free setting")
	rootCmd.Flags().String("min-free-check-interval", "", "Set the min-free-check-interval setting")
	rootCmd.Flags().String("nar-buffer-size", "", "Set the nar-buffer-size setting")
	rootCmd.Flags().String("narinfo-cache-negative-ttl", "", "Set the narinfo-cache-negative-ttl setting")
	rootCmd.Flags().String("narinfo-cache-positive-ttl", "", "Set the narinfo-cache-positive-ttl setting")
	rootCmd.Flags().String("netrc-file", "", "Set the netrc-file setting")
	rootCmd.Flags().String("nix-path", "", "Set the nix-path setting")
	rootCmd.Flags().Bool("no-accept-flake-config", false, "Disable the accept-flake-config setting")
	rootCmd.Flags().Bool("no-allow-dirty", false, "Disable the allow-dirty setting")
	rootCmd.Flags().Bool("no-allow-import-from-derivation", false, "Disable the allow-import-from-derivation setting")
	rootCmd.Flags().Bool("no-allow-new-privileges", false, "Disable the allow-new-privileges setting")
	rootCmd.Flags().Bool("no-allow-symlinked-store", false, "Disable the allow-symlinked-store setting")
	rootCmd.Flags().Bool("no-allow-unsafe-native-code-during-evaluation", false, "Disable the allow-unsafe-native-code-during-evaluation setting")
	rootCmd.Flags().Bool("no-auto-optimise-store", false, "Disable the auto-optimise-store setting")
	rootCmd.Flags().Bool("no-builders-use-substitutes", false, "Disable the builders-use-substitutes setting")
	rootCmd.Flags().Bool("no-compress-build-log", false, "Disable the compress-build-log setting")
	rootCmd.Flags().Bool("no-enforce-determinism", false, "Disable the enforce-determinism setting")
	rootCmd.Flags().Bool("no-eval-cache", false, "Disable the eval-cache setting")
	rootCmd.Flags().Bool("no-fallback", false, "Disable the fallback setting")
	rootCmd.Flags().Bool("no-filter-syscalls", false, "Disable the filter-syscalls setting")
	rootCmd.Flags().Bool("no-fsync-metadata", false, "Disable the fsync-metadata setting")
	rootCmd.Flags().Bool("no-http2", false, "Disable the http2 setting")
	rootCmd.Flags().Bool("no-ignore-try", false, "Disable the ignore-try setting")
	rootCmd.Flags().Bool("no-impersonate-linux-26", false, "Disable the impersonate-linux-26 setting")
	rootCmd.Flags().Bool("no-keep-build-log", false, "Disable the keep-build-log setting")
	rootCmd.Flags().Bool("no-keep-derivations", false, "Disable the keep-derivations setting")
	rootCmd.Flags().Bool("no-keep-env-derivations", false, "Disable the keep-env-derivations setting")
	rootCmd.Flags().Bool("no-keep-failed", false, "Disable the keep-failed setting")
	rootCmd.Flags().Bool("no-keep-going", false, "Disable the keep-going setting")
	rootCmd.Flags().Bool("no-keep-outputs", false, "Disable the keep-outputs setting")
	rootCmd.Flags().Bool("no-preallocate-contents", false, "Disable the preallocate-contents setting")
	rootCmd.Flags().Bool("no-print-missing", false, "Disable the print-missing setting")
	rootCmd.Flags().Bool("no-pure-eval", false, "Disable the pure-eval setting")
	rootCmd.Flags().Bool("no-require-sigs", false, "Disable the require-sigs setting")
	rootCmd.Flags().Bool("no-restrict-eval", false, "Disable the restrict-eval setting")
	rootCmd.Flags().Bool("no-run-diff-hook", false, "Disable the run-diff-hook setting")
	rootCmd.Flags().Bool("no-sandbox", false, "Disable sandboxing")
	rootCmd.Flags().Bool("no-sandbox-fallback", false, "Disable the sandbox-fallback setting")
	rootCmd.Flags().Bool("no-show-trace", false, "Disable the show-trace setting")
	rootCmd.Flags().Bool("no-substitute", false, "Disable the substitute setting")
	rootCmd.Flags().Bool("no-sync-before-registering", false, "Disable the sync-before-registering setting")
	rootCmd.Flags().Bool("no-trace-function-calls", false, "Disable the trace-function-calls setting")
	rootCmd.Flags().Bool("no-trace-verbose", false, "Disable the trace-verbose setting")
	rootCmd.Flags().Bool("no-use-case-hack", false, "Disable the use-case-hack setting")
	rootCmd.Flags().Bool("no-use-registries", false, "Disable the use-registries setting")
	rootCmd.Flags().Bool("no-use-sqlite-wal", false, "Disable the use-sqlite-wal setting")
	rootCmd.Flags().Bool("no-warn-dirty", false, "Disable the warn-dirty setting")
	rootCmd.Flags().String("plugin-files", "", "Set the plugin-files setting")
	rootCmd.Flags().String("post-build-hook", "", "Set the post-build-hook setting")
	rootCmd.Flags().String("pre-build-hook", "", "Set the pre-build-hook setting")
	rootCmd.Flags().Bool("preallocate-contents", false, "Enable the preallocate-contents setting")
	rootCmd.Flags().BoolP("print-build-logs", "L", false, "Print full build logs on standard error")
	rootCmd.Flags().Bool("print-missing", false, "Enable the print-missing setting")
	rootCmd.Flags().Bool("pure-eval", false, "Enable the pure-eval setting")
	rootCmd.Flags().Bool("quiet", false, "Decrease the logging verbosity level")
	rootCmd.Flags().Bool("relaxed-sandbox", false, "Enable sandboxing, but allow builds to disable it")
	rootCmd.Flags().String("repeat", "", "Set the repeat setting")
	rootCmd.Flags().Bool("require-sigs", false, "Enable the require-sigs setting")
	rootCmd.Flags().Bool("restrict-eval", false, "Enable the restrict-eval setting")
	rootCmd.Flags().Bool("run-diff-hook", false, "Enable the run-diff-hook setting")
	rootCmd.Flags().Bool("sandbox", false, "Enable sandboxing")
	rootCmd.Flags().String("sandbox-build-dir", "", "Set the sandbox-build-dir setting")
	rootCmd.Flags().String("sandbox-dev-shm-size", "", "Set the sandbox-dev-shm-size setting")
	rootCmd.Flags().Bool("sandbox-fallback", false, "Enable the sandbox-fallback setting")
	rootCmd.Flags().String("sandbox-paths", "", "Set the sandbox-paths setting")
	rootCmd.Flags().String("secret-key-files", "", "Set the secret-key-files setting")
	rootCmd.Flags().Bool("show-trace", false, "Enable the show-trace setting")
	rootCmd.Flags().String("stalled-download-timeout", "", "Set the stalled-download-timeout setting")
	rootCmd.Flags().String("store", "", "Set the store setting")
	rootCmd.Flags().Bool("substitute", false, "Enable the substitute setting")
	rootCmd.Flags().StringArray("substituters", nil, "Set the substituters setting")
	rootCmd.Flags().Bool("sync-before-registering", false, "Enable the sync-before-registering setting")
	rootCmd.Flags().String("system", "", "Set the system setting")
	rootCmd.Flags().String("system-features", "", "Set the system-features setting")
	rootCmd.Flags().String("tarball-ttl", "", "Set the tarball-ttl setting")
	rootCmd.Flags().String("timeout", "", "Set the timeout setting")
	rootCmd.Flags().Bool("trace-function-calls", false, "Enable the trace-function-calls setting")
	rootCmd.Flags().Bool("trace-verbose", false, "Enable the trace-verbose setting")
	rootCmd.Flags().String("trusted-public-keys", "", "Set the trusted-public-keys setting")
	rootCmd.Flags().String("trusted-substituters", "", "Set the trusted-substituters setting")
	rootCmd.Flags().String("trusted-users", "", "Set the trusted-users setting")
	rootCmd.Flags().Bool("use-case-hack", false, "Enable the use-case-hack setting")
	rootCmd.Flags().Bool("use-registries", false, "Enable the use-registries setting")
	rootCmd.Flags().Bool("use-sqlite-wal", false, "Enable the use-sqlite-wal setting")
	rootCmd.Flags().String("user-agent-suffix", "", "Set the user-agent-suffix setting")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase the logging verbosity level")
	rootCmd.Flags().Bool("warn-dirty", false, "Enable the warn-dirty setting")

	rootCmd.PersistentFlags().Bool("help", false, "Show usage information")
	rootCmd.PersistentFlags().Bool("offline", false, "Disable substituters and consider all previously downloaded files up-to-date")
	rootCmd.PersistentFlags().StringSlice("option", nil, "Set the Nix configuration setting name to value")
	rootCmd.PersistentFlags().Bool("refresh", false, "Consider all previously downloaded files out-of-date")
	rootCmd.PersistentFlags().Bool("version", false, "Show version information")

	rootCmd.Flag("option").Nargs = 2

	rootCmd.MarkFlagsMutuallyExclusive("accept-flake-config", "no-accept-flake-config")
	rootCmd.MarkFlagsMutuallyExclusive("allow-dirty", "no-allow-dirty")
	rootCmd.MarkFlagsMutuallyExclusive("allow-import-from-derivation", "no-allow-import-from-derivation")
	rootCmd.MarkFlagsMutuallyExclusive("allow-new-privileges", "no-allow-new-privileges")
	rootCmd.MarkFlagsMutuallyExclusive("allow-symlinked-store", "no-allow-symlinked-store")
	rootCmd.MarkFlagsMutuallyExclusive("allow-unsafe-native-code-during-evaluation", "no-allow-unsafe-native-code-during-evaluation")
	rootCmd.MarkFlagsMutuallyExclusive("auto-optimise-store", "no-auto-optimise-store")
	rootCmd.MarkFlagsMutuallyExclusive("builders-use-substitutes", "no-builders-use-substitutes")
	rootCmd.MarkFlagsMutuallyExclusive("compress-build-log", "no-compress-build-log")
	rootCmd.MarkFlagsMutuallyExclusive("enforce-determinism", "no-enforce-determinism")
	rootCmd.MarkFlagsMutuallyExclusive("eval-cache", "no-eval-cache")
	rootCmd.MarkFlagsMutuallyExclusive("fallback", "no-fallback")
	rootCmd.MarkFlagsMutuallyExclusive("filter-syscalls", "no-filter-syscalls")
	rootCmd.MarkFlagsMutuallyExclusive("fsync-metadata", "no-fsync-metadata")
	rootCmd.MarkFlagsMutuallyExclusive("http2", "no-http2")
	rootCmd.MarkFlagsMutuallyExclusive("ignore-try", "no-ignore-try")
	rootCmd.MarkFlagsMutuallyExclusive("impersonate-linux-26", "no-impersonate-linux-26")
	rootCmd.MarkFlagsMutuallyExclusive("keep-build-log", "no-keep-build-log")
	rootCmd.MarkFlagsMutuallyExclusive("keep-derivations", "no-keep-derivations")
	rootCmd.MarkFlagsMutuallyExclusive("keep-env-derivations", "no-keep-env-derivations")
	rootCmd.MarkFlagsMutuallyExclusive("keep-failed", "no-keep-failed")
	rootCmd.MarkFlagsMutuallyExclusive("keep-going", "no-keep-going")
	rootCmd.MarkFlagsMutuallyExclusive("keep-outputs", "no-keep-outputs")
	rootCmd.MarkFlagsMutuallyExclusive("preallocate-contents", "no-preallocate-contents")
	rootCmd.MarkFlagsMutuallyExclusive("print-missing", "no-print-missing")
	rootCmd.MarkFlagsMutuallyExclusive("pure-eval", "no-pure-eval")
	rootCmd.MarkFlagsMutuallyExclusive("require-sigs", "no-require-sigs")
	rootCmd.MarkFlagsMutuallyExclusive("restrict-eval", "no-restrict-eval")
	rootCmd.MarkFlagsMutuallyExclusive("run-diff-hook", "no-run-diff-hook")
	rootCmd.MarkFlagsMutuallyExclusive("sandbox", "no-sandbox")
	rootCmd.MarkFlagsMutuallyExclusive("sandbox-fallback", "no-sandbox-fallback")
	rootCmd.MarkFlagsMutuallyExclusive("show-trace", "no-show-trace")
	rootCmd.MarkFlagsMutuallyExclusive("substitute", "no-substitute")
	rootCmd.MarkFlagsMutuallyExclusive("sync-before-registering", "no-sync-before-registering")
	rootCmd.MarkFlagsMutuallyExclusive("trace-function-calls", "no-trace-function-calls")
	rootCmd.MarkFlagsMutuallyExclusive("trace-verbose", "no-trace-verbose")
	rootCmd.MarkFlagsMutuallyExclusive("use-case-hack", "no-use-case-hack")
	rootCmd.MarkFlagsMutuallyExclusive("use-registries", "no-use-registries")
	rootCmd.MarkFlagsMutuallyExclusive("use-sqlite-wal", "no-use-sqlite-wal")
	rootCmd.MarkFlagsMutuallyExclusive("warn-dirty", "no-warn-dirty")

	// TODO find out how flags shall be completed (nix doc insufficient)
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log-format": carapace.ActionValues("raw", "internal-json", "bar", "bar-with-logs"),
		"option": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return nix.ActionConfigKeys()
			case 1:
				return nix.ActionConfigValues(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
	})
}

func addEvaluationFlags(cmd *cobra.Command) {
	cmd.Flags().StringSlice("arg", nil, "Pass the value expr as the argument name to Nix functions")
	cmd.Flags().StringSlice("argstr", nil, "Pass the string string as the argument name to Nix functions")
	cmd.Flags().Bool("debugger", false, "Start an interactive environment if evaluation fail")
	cmd.Flags().String("eval-store", "", "The Nix store to use for evaluations")
	cmd.Flags().Bool("impure", false, "Allow access to mutable paths and repositories")
	cmd.Flags().BoolP("include", "I", false, "Add path to the list of locations used to look up <...> file names")
	cmd.Flags().String("override-flake", "", "Override the flake registries, redirecting original-ref to resolved-ref")

	cmd.Flag("arg").Nargs = 2
	cmd.Flag("argstr").Nargs = 2
}

func addFlakeFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("commit-lock-file", false, "Commit changes to the flake's lock file")
	cmd.Flags().String("inputs-from", "", "Use the inputs of the specified flake as registry entries")
	cmd.Flags().Bool("no-registries", false, "Don't allow lookups in the flake registries")
	cmd.Flags().Bool("no-update-lock-file", false, "Do not allow any updates to the flake's lock file")
	cmd.Flags().Bool("no-write-lock-file", false, "Do not write the flake's newly generated lock file")
	cmd.Flags().String("output-lock-file", "", "Write the given lock file instead of flake.lock")
	cmd.Flags().String("override-input", "", "Override a specific flake input (e.g. dwarffs/nixpkgs)")
	cmd.Flags().Bool("recreate-lock-file", false, "Recreate the flake's lock file from scratch")
	cmd.Flags().String("reference-lock-file", "", "Read the given lock file instead of flake.lock")
	cmd.Flags().String("update-input", "", "Update a specific flake input (ignoring its previous entry in the lock file")

	cmd.Flag("override-input").Nargs = 2
}

func addInterpretationFlags(cmd *cobra.Command) {
	cmd.Flags().String("expr", "", "Interpret installables as attribute paths relative to the Nix expression expr")
	cmd.Flags().StringP("file", "f", "", "Interpret installables as attribute paths relative to the Nix expression stored in file")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}

func addLoggingFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("debug", false, "Set the logging verbosity level to 'debug'")
	cmd.Flags().String("log-format", "", "Set the format of log output")
	cmd.Flags().BoolP("print-build-logs", "L", false, "Print full build logs on standard error")
	cmd.Flags().Bool("quiet", false, "Decrease the logging verbosity level")
	cmd.Flags().BoolP("verbose", "v", false, "Increase the logging verbosity level")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"log-format": carapace.ActionValues("raw", "internal-json", "bar", "bar-with-logs"),
	})
}
