package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/text"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var sendEmailCmd = &cobra.Command{
	Use:     "send-email",
	Short:   "Send a collection of patches as email",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {
	carapace.Gen(sendEmailCmd).Standalone()

	sendEmailCmd.Flags().String("8bit-encoding", "", "encoding to assume 8bit mails if undeclared")
	sendEmailCmd.Flags().Bool("annotate", false, "review each patch that will be sent in an editor")
	sendEmailCmd.Flags().String("batch-size", "", "send max <num> messages per connection")
	sendEmailCmd.Flags().StringSlice("bcc", nil, "email Bcc:")
	sendEmailCmd.Flags().StringSlice("cc", nil, "email Cc:")
	sendEmailCmd.Flags().String("cc-cmd", "", "email Cc: via `<str> $patch_path`.")
	sendEmailCmd.Flags().Bool("cc-cover", false, "email Cc: addresses in the cover letter.")
	sendEmailCmd.Flags().Bool("chain-reply-to", false, "send each email as a reply to the previous one")
	sendEmailCmd.Flags().Bool("compose", false, "open an editor for introduction")
	sendEmailCmd.Flags().String("compose-encoding", "", "encoding to assume for introduction")
	sendEmailCmd.Flags().String("confirm", "", "confirm recipients before sending")
	sendEmailCmd.Flags().Bool("dry-run", false, "don't actually send the emails")
	sendEmailCmd.Flags().Bool("dump-aliases", false, "dump configured aliases and exit")
	sendEmailCmd.Flags().String("envelope-sender", "", "email envelope sender")
	sendEmailCmd.Flags().Bool("force", false, "send even if safety checks would prevent it")
	sendEmailCmd.Flags().Bool("format-patch", false, "understand arguments as format-patch output")
	sendEmailCmd.Flags().String("from", "", "email From:")
	sendEmailCmd.Flags().String("header-cmd", "", "add headers via `<str> $patch_path`")
	sendEmailCmd.Flags().String("identity", "", "use the sendemail.<id> options")
	sendEmailCmd.Flags().String("in-reply-to", "", "email \"In-Reply-To:\"")
	sendEmailCmd.Flags().Bool("no-annotate", false, "do not review each patch in an editor")
	sendEmailCmd.Flags().Bool("no-bcc", false, "clear the Bcc: list")
	sendEmailCmd.Flags().Bool("no-cc", false, "clear the Cc: list")
	sendEmailCmd.Flags().Bool("no-cc-cover", false, "do not add Cc: addresses from the cover letter")
	sendEmailCmd.Flags().Bool("no-chain-reply-to", false, "send all emails after the first as replies to the first")
	sendEmailCmd.Flags().Bool("no-format-patch", false, "understand arguments as file names")
	sendEmailCmd.Flags().Bool("no-header-cmd", false, "disable any header command in use")
	sendEmailCmd.Flags().Bool("no-signed-off-by-cc", false, "do not add emails from Signed-off-by trailers to the cc list")
	sendEmailCmd.Flags().Bool("no-smtp-auth", false, "disable SMTP authentication")
	sendEmailCmd.Flags().Bool("no-suppress-from", false, "do not suppress the From: address from the cc list")
	sendEmailCmd.Flags().Bool("no-thread", false, "do not add In-Reply-To and References headers")
	sendEmailCmd.Flags().Bool("no-to", false, "clear the To: list")
	sendEmailCmd.Flags().Bool("no-to-cover", false, "do not add To: addresses from the cover letter")
	sendEmailCmd.Flags().Bool("no-validate", false, "do not perform patch sanity checks")
	sendEmailCmd.Flags().Bool("no-xmailer", false, "do not add the \"X-Mailer:\" header")
	sendEmailCmd.Flags().Bool("quiet", false, "output one line of info per email")
	sendEmailCmd.Flags().String("relogin-delay", "", "wait <int> seconds before reconnecting to SMTP server")
	sendEmailCmd.Flags().String("reply-to", "", "email \"Reply-To:\"")
	sendEmailCmd.Flags().String("sendmail-cmd", "", "command to run to send email")
	sendEmailCmd.Flags().Bool("signed-off-by-cc", false, "add emails from Signed-off-by trailers to the cc list")
	sendEmailCmd.Flags().String("smtp-auth", "", "space-separated list of allowed AUTH mechanisms")
	sendEmailCmd.Flags().String("smtp-debug", "", "enable or disable Net::SMTP debug")
	sendEmailCmd.Flags().String("smtp-domain", "", "the domain name sent to HELO/EHLO handshake")
	sendEmailCmd.Flags().String("smtp-encryption", "", "tls or ssl; anything else disables")
	sendEmailCmd.Flags().String("smtp-pass", "", "password for SMTP-AUTH")
	sendEmailCmd.Flags().String("smtp-server", "", "outgoing SMTP server to use")
	sendEmailCmd.Flags().StringSlice("smtp-server-option", nil, "outgoing SMTP server option to use")
	sendEmailCmd.Flags().String("smtp-server-port", "", "outgoing SMTP server port")
	sendEmailCmd.Flags().Bool("smtp-ssl", false, "deprecated alias for --smtp-encryption ssl")
	sendEmailCmd.Flags().String("smtp-ssl-cert-path", "", "path to ca-certificates (either directory or file)")
	sendEmailCmd.Flags().String("smtp-user", "", "username for SMTP-AUTH")
	sendEmailCmd.Flags().String("subject", "", "email \"Subject:\"")
	sendEmailCmd.Flags().String("suppress-cc", "", "specify category of recipients to suppress")
	sendEmailCmd.Flags().Bool("suppress-from", false, "suppress the From: address from the cc list")
	sendEmailCmd.Flags().Bool("thread", false, "add In-Reply-To and References headers")
	sendEmailCmd.Flags().StringArray("to", nil, "email To:")
	sendEmailCmd.Flags().String("to-cmd", "", "email To: via `<str> $patch_path`")
	sendEmailCmd.Flags().Bool("to-cover", false, "add To: addresses from the cover letter")
	sendEmailCmd.Flags().String("transfer-encoding", "", "transfer encoding to use (quoted-printable, 8bit, base64)")
	sendEmailCmd.Flags().Bool("validate", false, "perform patch sanity checks")
	sendEmailCmd.Flags().Bool("xmailer", false, "add the \"X-Mailer:\" header")
	rootCmd.AddCommand(sendEmailCmd)

	carapace.Gen(sendEmailCmd).FlagCompletion(carapace.ActionMap{
		"8bit-encoding":    text.ActionEncodings(),
		"cc-cmd":           bridge.ActionCarapaceBin().Split(),
		"compose-encoding": text.ActionEncodings(),
		"confirm": carapace.ActionValuesDescribed(
			"always", "will always confirm before sending",
			"never", "will never confirm before sending",
			"cc", "will confirm before sending when send-email has automatically added addresses from the patch to the Cc list",
			"compose", "will confirm before sending the first message when using --compose",
			"auto", "is equivalent to cc + compose",
		),
		"envelope-sender": carapace.ActionValues("auto"),
		"header-cmd":      bridge.ActionCarapaceBin().Split(),
		"sendmail-cmd":    bridge.ActionCarapaceBin().Split(),
		"smtp-auth":       carapace.ActionValues(), // TODO
		"smtp-debug": carapace.ActionValuesDescribed(
			"0", "disabled",
			"1", "enabled",
		).StyleF(style.ForKeyword),
		"smtp-encryption":    carapace.ActionValues("ssl", "tls"),
		"smtp-server-option": carapace.ActionValues(), // TODO
		"smtp-server-port":   net.ActionPorts(),
		"smtp-ssl-cert-path": carapace.ActionDirectories(),
		"to-cmd":             bridge.ActionCarapaceBin().Split(),
		"transfer-encoding":  carapace.ActionValues("7bit", "8bit", "quoted-printable", "base64", "auto"),
	})

	carapace.Gen(sendEmailCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
