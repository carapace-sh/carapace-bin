package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tsh",
	Short: "Teleport Command Line Client",
	Long:  "https://github.com/gravitational/teleport",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("add-keys-to-agent", "k", "", "Controls how keys are handled. Valid values are [auto no yes only].")
	rootCmd.PersistentFlags().String("auth", "", "Specify the name of authentication connector to use.")
	rootCmd.PersistentFlags().String("bind-addr", "", "Override host:port used when opening a browser for cluster logins")
	rootCmd.PersistentFlags().String("cert-format", "", "SSH certificate format")
	rootCmd.PersistentFlags().String("compat", "", "OpenSSH compatibility flag")
	rootCmd.PersistentFlags().Bool("completion-bash", false, "Output possible completions for the given args.")
	rootCmd.PersistentFlags().Bool("completion-script-bash", false, "Generate completion script for bash.")
	rootCmd.PersistentFlags().Bool("completion-script-zsh", false, "Generate completion script for ZSH.")
	rootCmd.PersistentFlags().String("cpu-profile", "", "Write CPU profile to file")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Verbose logging to stdout")
	rootCmd.PersistentFlags().Bool("enable-escape-sequences", false, "Enable support for SSH escape sequences. Type '~?' during an SSH session to list supported sequences. Default is enabled.")
	rootCmd.PersistentFlags().Bool("headless", false, "Use headless login. Shorthand for --auth=headless.")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show context-sensitive help (also try --help-long and --help-man).")
	rootCmd.PersistentFlags().Bool("help-long", false, "Generate long help.")
	rootCmd.PersistentFlags().Bool("help-man", false, "Generate a man page.")
	rootCmd.PersistentFlags().StringP("identity", "i", "", "Identity file")
	rootCmd.PersistentFlags().Bool("insecure", false, "Do not verify server's certificate and host name. Use only in test environments")
	rootCmd.PersistentFlags().StringP("jumphost", "J", "", "SSH jumphost")
	rootCmd.PersistentFlags().StringP("login", "l", "", "Remote host login")
	rootCmd.PersistentFlags().String("mem-profile", "", "Write memory profile to file")
	rootCmd.PersistentFlags().String("mfa-mode", "", "Preferred mode for MFA and Passwordless assertions (auto, cross-platform, platform, otp)")
	rootCmd.PersistentFlags().String("mlock", "", "Determines whether process memory will be locked and whether failure to do so will be accepted (off, auto, best_effort, strict).")
	rootCmd.PersistentFlags().String("namespace", "", "Namespace of the cluster")
	rootCmd.PersistentFlags().Bool("no-completion-bash", false, "Output possible completions for the given args.")
	rootCmd.PersistentFlags().Bool("no-completion-script-bash", false, "Generate completion script for bash.")
	rootCmd.PersistentFlags().Bool("no-completion-script-zsh", false, "Generate completion script for ZSH.")
	rootCmd.PersistentFlags().Bool("no-debug", false, "Verbose logging to stdout")
	rootCmd.PersistentFlags().Bool("no-enable-escape-sequences", false, "Enable support for SSH escape sequences. Type '~?' during an SSH session to list supported sequences. Default is enabled.")
	rootCmd.PersistentFlags().Bool("no-headless", false, "Use headless login. Shorthand for --auth=headless.")
	rootCmd.PersistentFlags().Bool("no-help", false, "Show context-sensitive help (also try --help-long and --help-man).")
	rootCmd.PersistentFlags().Bool("no-help-long", false, "Generate long help.")
	rootCmd.PersistentFlags().Bool("no-help-man", false, "Generate a man page.")
	rootCmd.PersistentFlags().Bool("no-insecure", false, "Do not verify server's certificate and host name. Use only in test environments")
	rootCmd.PersistentFlags().Bool("no-nocache", false, "do not cache cluster discovery locally")
	rootCmd.PersistentFlags().Bool("no-skip-version-check", false, "Skip version checking between server and client.")
	rootCmd.PersistentFlags().Bool("no-trace", false, "Capture and export distributed traces")
	rootCmd.PersistentFlags().Bool("no-use-local-ssh-agent", false, "Deprecated in favor of the add-keys-to-agent flag.")
	rootCmd.PersistentFlags().Bool("nocache", false, "do not cache cluster discovery locally")
	rootCmd.PersistentFlags().StringP("option", "o", "", "")
	rootCmd.PersistentFlags().String("proxy", "", "Teleport proxy address")
	rootCmd.PersistentFlags().Bool("skip-version-check", false, "Skip version checking between server and client.")
	rootCmd.PersistentFlags().Bool("trace", false, "Capture and export distributed traces")
	rootCmd.PersistentFlags().String("trace-exporter", "", "An OTLP exporter URL to send spans to. Note - only tsh spans will be included.")
	rootCmd.PersistentFlags().String("trace-profile", "", "Write trace profile to file")
	rootCmd.PersistentFlags().String("ttl", "", "Minutes to live for a session")
	rootCmd.PersistentFlags().Bool("use-local-ssh-agent", false, "Deprecated in favor of the add-keys-to-agent flag.")
	rootCmd.PersistentFlags().String("user", "", "Teleport user, defaults to current local user")
	rootCmd.Flag("compat").Hidden = true
	rootCmd.Flag("completion-bash").Hidden = true
	rootCmd.Flag("completion-script-bash").Hidden = true
	rootCmd.Flag("completion-script-zsh").Hidden = true
	rootCmd.Flag("cpu-profile").Hidden = true
	rootCmd.Flag("help").Hidden = true
	rootCmd.Flag("help-long").Hidden = true
	rootCmd.Flag("help-man").Hidden = true
	rootCmd.Flag("mem-profile").Hidden = true
	rootCmd.Flag("namespace").Hidden = true
	rootCmd.Flag("no-completion-bash").Hidden = true
	rootCmd.Flag("no-completion-script-bash").Hidden = true
	rootCmd.Flag("no-completion-script-zsh").Hidden = true
	rootCmd.Flag("no-debug").Hidden = true
	rootCmd.Flag("no-enable-escape-sequences").Hidden = true
	rootCmd.Flag("no-headless").Hidden = true
	rootCmd.Flag("no-help").Hidden = true
	rootCmd.Flag("no-help-long").Hidden = true
	rootCmd.Flag("no-help-man").Hidden = true
	rootCmd.Flag("no-insecure").Hidden = true
	rootCmd.Flag("no-nocache").Hidden = true
	rootCmd.Flag("no-skip-version-check").Hidden = true
	rootCmd.Flag("no-trace").Hidden = true
	rootCmd.Flag("no-use-local-ssh-agent").Hidden = true
	rootCmd.Flag("nocache").Hidden = true
	rootCmd.Flag("option").Hidden = true
	rootCmd.Flag("trace").Hidden = true
	rootCmd.Flag("trace-exporter").Hidden = true
	rootCmd.Flag("trace-profile").Hidden = true
	rootCmd.Flag("use-local-ssh-agent").Hidden = true

	// TODO missing flags
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add-keys-to-agent": carapace.ActionValues("auto", "no", "yes", "only").StyleF(style.ForKeyword),
		"cpu-profile":       carapace.ActionFiles(),
		"identity": carapace.Batch(
			ssh.ActionPrivateKeys(),
			carapace.ActionFiles(),
		).ToA(),
		"mem-profile":   carapace.ActionFiles(),
		"mfa-mode":      carapace.ActionValues("auto", "cross-platform", "platform", "otp").StyleF(style.ForKeyword),
		"mlock":         carapace.ActionValues("off", "auto", "best_effort", "strict").StyleF(style.ForKeyword),
		"trace-profile": carapace.ActionFiles(),
	})
}
