package common

import (
	"github.com/spf13/cobra"
)

func AddMiscellaneousFlags(cmd *cobra.Command) {
	cmd.Flags().StringSlice("build_metadata", []string{}, "Custom key-value string pairs to supply in a build event")
	cmd.Flags().String("color", "", "Use terminal controls to colorize output")
	cmd.Flags().StringSlice("config", []string{}, "Selects additional config sections from the rc files; for every <command>, it also pulls in the options from <command>:<config> if such a section exists; if this section doesn't exist in any .rc file, Blaze fails with an error")
	cmd.Flags().StringSlice("credential_helper", []string{}, "Configures a credential helper conforming to the <a href=\"https://github")
	cmd.Flags().String("credential_helper_cache_duration", "", "The default duration for which credentials supplied by a credential helper are cached if the helper does not provide when the credentials expire")
	cmd.Flags().String("credential_helper_timeout", "", "Configures the timeout for a credential helper")
	cmd.Flags().String("curses", "", "Use terminal cursor controls to minimize scrolling output")
	cmd.Flags().String("disk_cache", "", "A path to a directory where Bazel can read and write actions and action outputs")
	cmd.Flags().Bool("enable_platform_specific_config", false, "If true, Bazel picks up host-OS-specific config lines from bazelrc files")
	cmd.Flags().Bool("experimental_rule_extension_api", false, "Enable experimental rule extension API and subrule APIs Tags: loading_and_analysis, experimental")
	cmd.Flags().Bool("experimental_windows_watchfs", false, "If true, experimental Windows support for --watchfs is enabled")
	cmd.Flags().String("google_auth_scopes", "", "A comma-separated list of Google Cloud authentication scopes")
	cmd.Flags().String("google_credentials", "", "Specifies the file to get authentication credentials from")
	cmd.Flags().Bool("google_default_credentials", false, "Whether to use 'Google Application Default Credentials' for authentication")
	cmd.Flags().String("grpc_keepalive_time", "", "Configures keep-alive pings for outgoing gRPC connections")
	cmd.Flags().String("grpc_keepalive_timeout", "", "Configures a keep-alive timeout for outgoing gRPC connections")
	cmd.Flags().Bool("incompatible_disable_non_executable_java_binary", false, "If true, java_binary is always executable")
	cmd.Flags().Bool("incompatible_disallow_symlink_file_to_dir", false, "No-op")
	cmd.Flags().Bool("no-enable_platform_specific_config", false, "If true, Bazel picks up host-OS-specific config lines from bazelrc files")
	cmd.Flags().Bool("no-experimental_rule_extension_api", false, "Enable experimental rule extension API and subrule APIs Tags: loading_and_analysis, experimental")
	cmd.Flags().Bool("no-experimental_windows_watchfs", false, "If true, experimental Windows support for --watchfs is enabled")
	cmd.Flags().Bool("no-google_default_credentials", false, "Whether to use 'Google Application Default Credentials' for authentication")
	cmd.Flags().Bool("no-incompatible_disable_non_executable_java_binary", false, "If true, java_binary is always executable")
	cmd.Flags().Bool("no-incompatible_disallow_symlink_file_to_dir", false, "No-op")
	cmd.Flags().Bool("no-progress_in_terminal_title", false, "Show the command progress in the terminal title")
	cmd.Flags().Bool("no-show_progress", false, "Display progress messages during a build")
	cmd.Flags().Bool("no-show_timestamps", false, "Include timestamps in messages")
	cmd.Flags().Bool("no-watchfs", false, "On Linux/macOS: If true, bazel tries to use the operating system's file watch service for local changes instead of scanning every file for a change")
	cmd.Flags().StringSlice("override_repository", []string{}, "Override a repository with a local path in the form of <repository name>=<path>")
	cmd.Flags().Bool("progress_in_terminal_title", false, "Show the command progress in the terminal title")
	cmd.Flags().Bool("show_progress", false, "Display progress messages during a build")
	cmd.Flags().String("show_progress_rate_limit", "", "Minimum number of seconds between progress messages in the output")
	cmd.Flags().Bool("show_timestamps", false, "Include timestamps in messages")
	cmd.Flags().String("tls_certificate", "", "Specify a path to a TLS certificate that is trusted to sign server certificates")
	cmd.Flags().String("tls_client_certificate", "", "Specify the TLS client certificate to use; you also need to provide a client key to enable client authentication")
	cmd.Flags().String("tls_client_key", "", "Specify the TLS client key to use; you also need to provide a client certificate to enable client authentication")
	cmd.Flags().String("ui_actions_shown", "", "Number of concurrent actions shown in the detailed progress bar; each action is shown on a separate line")
	cmd.Flags().Bool("watchfs", false, "On Linux/macOS: If true, bazel tries to use the operating system's file watch service for local changes instead of scanning every file for a change")

	// TODO add flag completion
}
