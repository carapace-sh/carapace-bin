package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tsh",
	Short: "Teleport Command Line Client.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("add-keys-to-agent", "k", "auto", "Controls how keys are handled. Valid values are [auto no yes only].")
	rootCmd.PersistentFlags().String("auth", "", "Specify the name of authentication connector to use.")
	rootCmd.PersistentFlags().String("bind-addr", "", "Override host:port used when opening a browser for cluster logins.")
	rootCmd.PersistentFlags().String("browser-login", "", "Set to 'none' to suppress browser opening on login.")
	rootCmd.PersistentFlags().String("callback", "", "Override the base URL (host:port) of the link shown when opening a browser for cluster logins. Must be used with --bind-addr.")
	rootCmd.PersistentFlags().String("cert-format", "", "SSH certificate format.")
	rootCmd.PersistentFlags().Bool("check-update", false, "Check for availability of managed update.")
	rootCmd.PersistentFlags().String("compat", "", "OpenSSH compatibility flag.")
	rootCmd.PersistentFlags().Bool("completion-bash", false, "Output possible completions for the given args.")
	rootCmd.PersistentFlags().Bool("completion-script-bash", false, "Generate completion script for bash.")
	rootCmd.PersistentFlags().Bool("completion-script-fish", false, "Generate completion script for fish.")
	rootCmd.PersistentFlags().Bool("completion-script-zsh", false, "Generate completion script for ZSH.")
	rootCmd.PersistentFlags().String("cpu-profile", "", "Write CPU profile to file.")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Verbose logging to stdout.")
	rootCmd.PersistentFlags().Bool("enable-escape-sequences", true, "Enable support for SSH escape sequences. Type '~?' during an SSH session to list supported sequences. Default is enabled.")
	rootCmd.PersistentFlags().String("fork-kill-fd", "", "File descriptor to check parent health on when forked. For internal use only.")
	rootCmd.PersistentFlags().String("fork-signal-fd", "", "File descriptor to signal parent on when forked. Overrides --fork-after-authentication. For internal use only.")
	rootCmd.PersistentFlags().Bool("headless", false, "Use headless login. Shorthand for --auth=headless.")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show context-sensitive help (also try --help-long and --help-man).")
	rootCmd.PersistentFlags().Bool("help-long", false, "Generate long help.")
	rootCmd.PersistentFlags().Bool("help-man", false, "Generate a man page.")
	rootCmd.PersistentFlags().StringP("identity", "i", "", "Identity file.")
	rootCmd.PersistentFlags().Bool("insecure", false, "Do not verify server's certificate and host name. Use only in test environments.")
	rootCmd.PersistentFlags().StringP("jumphost", "J", "", "SSH jumphost.")
	rootCmd.PersistentFlags().StringP("login", "l", "", "Remote host login.")
	rootCmd.PersistentFlags().String("mem-profile", "", "Write memory profile to file.")
	rootCmd.PersistentFlags().String("mfa-mode", "auto", "Preferred mode for MFA and Passwordless assertions (auto, cross-platform, platform, otp, sso, browser).")
	rootCmd.PersistentFlags().String("mlock", "auto", "Determines whether process memory will be locked and whether failure to do so will be accepted (off, auto, best_effort, strict).")
	rootCmd.PersistentFlags().Bool("no-check-update", false, "Check for availability of managed update.")
	rootCmd.PersistentFlags().Bool("no-completion-bash", false, "Output possible completions for the given args.")
	rootCmd.PersistentFlags().Bool("no-completion-script-bash", false, "Generate completion script for bash.")
	rootCmd.PersistentFlags().Bool("no-completion-script-fish", false, "Generate completion script for fish.")
	rootCmd.PersistentFlags().Bool("no-completion-script-zsh", false, "Generate completion script for ZSH.")
	rootCmd.PersistentFlags().Bool("no-debug", false, "Verbose logging to stdout.")
	rootCmd.PersistentFlags().Bool("no-enable-escape-sequences", false, "Enable support for SSH escape sequences. Type '~?' during an SSH session to list supported sequences. Default is enabled.")
	rootCmd.PersistentFlags().Bool("no-headless", false, "Use headless login. Shorthand for --auth=headless.")
	rootCmd.PersistentFlags().Bool("no-help", false, "Show context-sensitive help (also try --help-long and --help-man).")
	rootCmd.PersistentFlags().Bool("no-help-long", false, "Generate long help.")
	rootCmd.PersistentFlags().Bool("no-help-man", false, "Generate a man page.")
	rootCmd.PersistentFlags().Bool("no-insecure", false, "Do not verify server's certificate and host name. Use only in test environments.")
	rootCmd.PersistentFlags().Bool("no-nocache", false, "Do not cache cluster discovery locally.")
	rootCmd.PersistentFlags().Bool("no-os-log", false, "Verbose logging to the unified logging system. This flag implies --debug. Also available through the TELEPORT_OS_LOG env var. More details see https://goteleport.com/docs/connect-your-client/tsh/#debug-logs.")
	rootCmd.PersistentFlags().Bool("no-skip-version-check", false, "Skip version checking between server and client.")
	rootCmd.PersistentFlags().Bool("no-trace", false, "Capture and export distributed traces.")
	rootCmd.PersistentFlags().Bool("no-use-local-ssh-agent", false, "Deprecated in favor of the add-keys-to-agent flag.")
	rootCmd.PersistentFlags().Bool("nocache", false, "Do not cache cluster discovery locally.")
	rootCmd.PersistentFlags().StringP("option", "o", "", "")
	rootCmd.PersistentFlags().Bool("os-log", false, "Verbose logging to the unified logging system. This flag implies --debug. Also available through the TELEPORT_OS_LOG env var. More details see https://goteleport.com/docs/connect-your-client/tsh/#debug-logs.")
	rootCmd.PersistentFlags().String("piv-slot", "", "Specify a PIV slot key to use for Hardware Key support instead of the default. Ex: \"9d\".")
	rootCmd.PersistentFlags().String("proxy", "", "Teleport proxy address.")
	rootCmd.PersistentFlags().String("relay", "", "Teleport relay address, \"none\" to explicitly disable the use of a relay, or \"default\" to use the cluster-provided address even if a different address was specified at login time.")
	rootCmd.PersistentFlags().Bool("skip-version-check", false, "Skip version checking between server and client.")
	rootCmd.PersistentFlags().Bool("trace", false, "Capture and export distributed traces.")
	rootCmd.PersistentFlags().String("trace-exporter", "", "An OTLP exporter URL to send spans to. Note - only tsh spans will be included.")
	rootCmd.PersistentFlags().String("trace-profile", "", "Write trace profile to file.")
	rootCmd.PersistentFlags().String("ttl", "", "Minutes to live for a session.")
	rootCmd.PersistentFlags().Bool("use-local-ssh-agent", true, "Deprecated in favor of the add-keys-to-agent flag.")
	rootCmd.PersistentFlags().String("user", "", "Teleport user, defaults to current local user.")
	rootCmd.Flag("browser-login").Hidden = true
	rootCmd.Flag("check-update").Hidden = true
	rootCmd.Flag("compat").Hidden = true
	rootCmd.Flag("completion-bash").Hidden = true
	rootCmd.Flag("completion-script-bash").Hidden = true
	rootCmd.Flag("completion-script-fish").Hidden = true
	rootCmd.Flag("completion-script-zsh").Hidden = true
	rootCmd.Flag("cpu-profile").Hidden = true
	rootCmd.Flag("fork-kill-fd").Hidden = true
	rootCmd.Flag("fork-signal-fd").Hidden = true
	rootCmd.Flag("help").Hidden = true
	rootCmd.Flag("help-long").Hidden = true
	rootCmd.Flag("help-man").Hidden = true
	rootCmd.Flag("mem-profile").Hidden = true
	rootCmd.Flag("no-check-update").Hidden = true
	rootCmd.Flag("no-completion-bash").Hidden = true
	rootCmd.Flag("no-completion-script-bash").Hidden = true
	rootCmd.Flag("no-completion-script-fish").Hidden = true
	rootCmd.Flag("no-completion-script-zsh").Hidden = true
	rootCmd.Flag("no-debug").Hidden = true
	rootCmd.Flag("no-enable-escape-sequences").Hidden = true
	rootCmd.Flag("no-headless").Hidden = true
	rootCmd.Flag("no-help").Hidden = true
	rootCmd.Flag("no-help-long").Hidden = true
	rootCmd.Flag("no-help-man").Hidden = true
	rootCmd.Flag("no-insecure").Hidden = true
	rootCmd.Flag("no-nocache").Hidden = true
	rootCmd.Flag("no-os-log").Hidden = true
	rootCmd.Flag("no-skip-version-check").Hidden = true
	rootCmd.Flag("no-trace").Hidden = true
	rootCmd.Flag("no-use-local-ssh-agent").Hidden = true
	rootCmd.Flag("nocache").Hidden = true
	rootCmd.Flag("option").Hidden = true
	rootCmd.Flag("os-log").Hidden = true
	rootCmd.Flag("trace").Hidden = true
	rootCmd.Flag("trace-exporter").Hidden = true
	rootCmd.Flag("trace-profile").Hidden = true
	rootCmd.Flag("use-local-ssh-agent").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add-keys-to-agent": carapace.ActionValues("auto", "no", "yes", "only").StyleF(style.ForKeyword),
		"cpu-profile":       carapace.ActionFiles(),
		"identity": carapace.Batch(
			ssh.ActionPrivateKeys(),
			carapace.ActionFiles(),
		).ToA(),
		"mem-profile":   carapace.ActionFiles(),
		"mfa-mode":      carapace.ActionValues("auto", "cross-platform", "platform", "otp", "sso", "browser").StyleF(style.ForKeyword),
		"mlock":         carapace.ActionValues("off", "auto", "best_effort", "strict").StyleF(style.ForKeyword),
		"trace-profile": carapace.ActionFiles(),
	})
}
