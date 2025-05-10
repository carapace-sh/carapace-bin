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

	sendEmailCmd.Flags().String("8bit-encoding", "", "Encoding to assume 8bit mails if undeclared")
	sendEmailCmd.Flags().Bool("annotate", false, "Review each patch that will be sent in an editor.")
	sendEmailCmd.Flags().String("batch-size", "", "send max  message per connection.")
	sendEmailCmd.Flags().StringSlice("bcc", nil, "Email Bcc:")
	sendEmailCmd.Flags().StringSlice("cc", nil, "Email Cc:")
	sendEmailCmd.Flags().String("cc-cmd", "", "Email Cc: via `<str> $patch_path`.")
	sendEmailCmd.Flags().Bool("cc-cover", false, "Email Cc: addresses in the cover letter.")
	sendEmailCmd.Flags().Bool("chain-reply-to", false, "Chain In-Reply-To: fields. Default off.")
	sendEmailCmd.Flags().Bool("compose", false, "Open an editor for introduction.")
	sendEmailCmd.Flags().String("compose-encoding", "", "Encoding to assume for introduction.")
	sendEmailCmd.Flags().String("confirm", "", "Confirm recipients before sending;")
	sendEmailCmd.Flags().Bool("dry-run", false, "Don't actually send the emails.")
	sendEmailCmd.Flags().Bool("dump-aliases", false, "Dump configured aliases and exit.")
	sendEmailCmd.Flags().String("envelope-sender", "", "Email envelope sender.")
	sendEmailCmd.Flags().Bool("force", false, "Send even if safety checks would prevent it.")
	sendEmailCmd.Flags().Bool("format-patch", false, "understand any non optional arguments as")
	sendEmailCmd.Flags().String("from", "", "Email From:")
	sendEmailCmd.Flags().String("header-cmd", "", "Add headers via `<str> $patch_path`.")
	sendEmailCmd.Flags().String("identity", "", "Use the sendemail.<id> options.")
	sendEmailCmd.Flags().String("in-reply-to", "", "Email \"In-Reply-To:\"")
	sendEmailCmd.Flags().Bool("no-annotate", false, "Review each patch that will be sent in an editor.")
	sendEmailCmd.Flags().Bool("no-bcc", false, "Email Bcc:")
	sendEmailCmd.Flags().Bool("no-cc", false, "Email Cc:")
	sendEmailCmd.Flags().Bool("no-cc-cover", false, "Email Cc: addresses in the cover letter.")
	sendEmailCmd.Flags().Bool("no-chain-reply-to", false, "Chain In-Reply-To: fields. Default off.")
	sendEmailCmd.Flags().Bool("no-format-patch", false, "understand any non optional arguments as")
	sendEmailCmd.Flags().Bool("no-header-cmd", false, "Disable any header command in use.")
	sendEmailCmd.Flags().Bool("no-signed-off-by-cc", false, "Send to Signed-off-by: addresses. Default on.")
	sendEmailCmd.Flags().Bool("no-smtp-auth", false, "Disable SMTP authentication. Shorthand for")
	sendEmailCmd.Flags().Bool("no-suppress-from", false, "Send to self. Default off.")
	sendEmailCmd.Flags().Bool("no-thread", false, "Use In-Reply-To: field. Default on.")
	sendEmailCmd.Flags().Bool("no-to", false, "Email To:")
	sendEmailCmd.Flags().Bool("no-to-cover", false, "Email To: addresses in the cover letter.")
	sendEmailCmd.Flags().Bool("no-validate", false, "Perform patch sanity checks. Default on.")
	sendEmailCmd.Flags().Bool("no-xmailer", false, "Add \"X-Mailer:\" header (default).")
	sendEmailCmd.Flags().Bool("quiet", false, "Output one line of info per email.")
	sendEmailCmd.Flags().String("relogin-delay", "", "delay seconds between two successive login.")
	sendEmailCmd.Flags().String("reply-to", "", "Email \"Reply-To:\"")
	sendEmailCmd.Flags().String("sendmail-cmd", "", "Command to run to send email.")
	sendEmailCmd.Flags().Bool("signed-off-by-cc", false, "Send to Signed-off-by: addresses. Default on.")
	sendEmailCmd.Flags().String("smtp-auth", "", "Space-separated list of allowed AUTH mechanisms, or")
	sendEmailCmd.Flags().String("smtp-debug", "", "Disable, enable Net::SMTP debug.")
	sendEmailCmd.Flags().String("smtp-domain", "", "The domain name sent to HELO/EHLO handshake")
	sendEmailCmd.Flags().String("smtp-encryption", "", "tls or ssl; anything else disables.")
	sendEmailCmd.Flags().String("smtp-pass", "", "Password for SMTP-AUTH; not necessary.")
	sendEmailCmd.Flags().String("smtp-server", "", "Outgoing SMTP server to use. The port")
	sendEmailCmd.Flags().StringSlice("smtp-server-option", nil, "Outgoing SMTP server option to use.")
	sendEmailCmd.Flags().String("smtp-server-port", "", "Outgoing SMTP server port.")
	sendEmailCmd.Flags().Bool("smtp-ssl", false, "Deprecated. Use '--smtp-encryption ssl'.")
	sendEmailCmd.Flags().String("smtp-ssl-cert-path", "", "Path to ca-certificates (either directory or file).")
	sendEmailCmd.Flags().String("smtp-user", "", "Username for SMTP-AUTH.")
	sendEmailCmd.Flags().String("subject", "", "Email \"Subject:\"")
	sendEmailCmd.Flags().String("suppress-cc", "", "author, self, sob, cc, cccmd, body, bodycc, misc-by, all.")
	sendEmailCmd.Flags().Bool("suppress-from", false, "Send to self. Default off.")
	sendEmailCmd.Flags().Bool("thread", false, "Use In-Reply-To: field. Default on.")
	sendEmailCmd.Flags().String("to", "", "Email To:")
	sendEmailCmd.Flags().String("to-cmd", "", "Email To: via `<str> $patch_path`.")
	sendEmailCmd.Flags().Bool("to-cover", false, "Email To: addresses in the cover letter.")
	sendEmailCmd.Flags().String("transfer-encoding", "", "Transfer encoding to use (quoted-printable, 8bit, base64)")
	sendEmailCmd.Flags().Bool("validate", false, "Perform patch sanity checks. Default on.")
	sendEmailCmd.Flags().Bool("xmailer", false, "Add \"X-Mailer:\" header (default).")
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
