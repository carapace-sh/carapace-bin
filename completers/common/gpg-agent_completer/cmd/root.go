package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gpg-agent",
	Short: "Secret key management for GnuPG",
	Long:  "https://linux.die.net/man/1/gpg-agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("allow-emacs-pinentry", false, "allow passphrase to be prompted through Emacs")
	rootCmd.Flags().Bool("allow-preset-passphrase", false, "allow presetting passphrase")
	rootCmd.Flags().Bool("batch", false, "never use the PIN-entry")
	rootCmd.Flags().String("check-passphrase-pattern", "", "check new passphrases against pattern in FILE")
	rootCmd.Flags().BoolP("csh", "c", false, "csh-style command output")
	rootCmd.Flags().Bool("daemon", false, "run in daemon mode (background)")
	rootCmd.Flags().String("default-cache-ttl", "", "expire cached PINs after N seconds")
	rootCmd.Flags().String("default-cache-ttl-ssh", "", "expire SSH keys after N seconds")
	rootCmd.Flags().Bool("disable-scdaemon", false, "do not use the SCdaemon")
	rootCmd.Flags().Bool("enable-passphrase-history", false, "do not allow the reuse of old passphrases")
	rootCmd.Flags().Bool("enable-ssh-support", false, "enable ssh support")
	rootCmd.Flags().Bool("enforce-passphrase-constraints", false, "do not allow bypassing the passphrase policy")
	rootCmd.Flags().String("extra-socket", "", "accept some commands via NAME")
	rootCmd.Flags().Bool("grab", false, "let PIN-Entry grab keyboard and mouse")
	rootCmd.Flags().Bool("ignore-cache-for-signing", false, "do not use the PIN cache when signing")
	rootCmd.Flags().Bool("keep-display", false, "ignore requests to change the X display")
	rootCmd.Flags().Bool("keep-tty", false, "ignore requests to change the TTY")
	rootCmd.Flags().String("log-file", "", "write server mode logs to FILE")
	rootCmd.Flags().String("max-cache-ttl", "", "set maximum PIN cache lifetime to N seconds")
	rootCmd.Flags().String("max-cache-ttl-ssh", "", "set maximum SSH key lifetime to N seconds")
	rootCmd.Flags().String("max-passphrase-days", "", "expire the passphrase after N days")
	rootCmd.Flags().String("min-passphrase-len", "", "set minimal required length for new passphrases to N")
	rootCmd.Flags().String("min-passphrase-nonalpha", "", "require at least N non-alpha characters for a new passphrase")
	rootCmd.Flags().Bool("no-allow-external-cache", false, "disallow the use of an external password cache")
	rootCmd.Flags().Bool("no-allow-loopback-pinentry", false, "disallow caller to override the pinentry")
	rootCmd.Flags().Bool("no-allow-mark-trusted", false, "disallow clients to mark keys as \"trusted\"")
	rootCmd.Flags().Bool("no-detach", false, "do not detach from the console")
	rootCmd.Flags().String("options", "", "read options from FILE")
	rootCmd.Flags().String("pinentry-program", "", "use PGM as the PIN-Entry program")
	rootCmd.Flags().String("pinentry-timeout", "", "set the Pinentry timeout to N seconds")
	rootCmd.Flags().BoolP("quiet", "q", false, "be somewhat more quiet")
	rootCmd.Flags().String("scdaemon-program", "", "use PGM as the SCdaemon program")
	rootCmd.Flags().Bool("server", false, "run in server mode (foreground)")
	rootCmd.Flags().BoolP("sh", "s", false, "sh-style command output")
	rootCmd.Flags().String("ssh-fingerprint-digest", "", "use ALGO to show ssh fingerprints")
	rootCmd.Flags().Bool("supervised", false, "run in supervised mode")
	rootCmd.Flags().String("tpm2daemon-program", "", "use PGM as the tpm2daemon program")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"check-passphrase-pattern": carapace.ActionFiles(),
		"log-file":                 carapace.ActionFiles(),
		"options":                  carapace.ActionFiles(),
		"pinentry-program":         carapace.ActionFiles(),
		"scdaemon-program":         carapace.ActionFiles(),
		"tpm2daemon-program":       carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("daemon").Changed {
				return carapace.ActionFiles()
			}
			return carapace.ActionValues()
		}),
	)
}
