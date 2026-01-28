package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var smimeCmd = &cobra.Command{
	Use:     "smime",
	Short:   "S/MIME mail processing",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(smimeCmd).Standalone()

	smimeCmd.Flags().StringS("CAfile", "CAfile", "", "Trusted certificates file")
	smimeCmd.Flags().StringS("CApath", "CApath", "", "Trusted certificates directory")
	smimeCmd.Flags().StringS("CAstore", "CAstore", "", "Trusted certificates store URI")
	smimeCmd.Flags().BoolS("binary", "binary", false, "Don't translate message to text")
	smimeCmd.Flags().StringS("certfile", "certfile", "", "Extra signer and intermediate CA certificates to include when signing or to use as preferred signer certs and for chain building when verifying")
	smimeCmd.Flags().StringS("config", "config", "", "Load a configuration file (this may load modules)")
	smimeCmd.Flags().StringS("content", "content", "", "Supply or override content for detached signature")
	smimeCmd.Flags().BoolS("crlfeol", "crlfeol", false, "Use CRLF as EOL termination instead of LF only")
	smimeCmd.Flags().BoolS("decrypt", "decrypt", false, "Decrypt encrypted message")
	smimeCmd.Flags().BoolS("encrypt", "encrypt", false, "Encrypt message")
	smimeCmd.Flags().StringS("from", "from", "", "From address")
	smimeCmd.Flags().StringS("in", "in", "", "Input file")
	smimeCmd.Flags().BoolS("indef", "indef", false, "Same as -stream")
	smimeCmd.Flags().StringS("inform", "inform", "", "Input format SMIME (default), PEM or DER")
	smimeCmd.Flags().StringS("inkey", "inkey", "", "Input private key (if not signer or recipient)")
	smimeCmd.Flags().StringS("keyform", "keyform", "", "Input private key format (DER/PEM)")
	smimeCmd.Flags().StringS("md", "md", "", "Digest algorithm to use when signing or resigning")
	smimeCmd.Flags().BoolS("no-CAfile", "no-CAfile", false, "Do not load the default certificates file")
	smimeCmd.Flags().BoolS("no-CApath", "no-CApath", false, "Do not load certificates from the default certificates directory")
	smimeCmd.Flags().BoolS("no-CAstore", "no-CAstore", false, "Do not load certificates from the default certificates store")
	smimeCmd.Flags().BoolS("noattr", "noattr", false, "Don't include any signed attributes")
	smimeCmd.Flags().BoolS("nocerts", "nocerts", false, "Don't include signers certificate when signing")
	smimeCmd.Flags().BoolS("nochain", "nochain", false, "set PKCS7_NOCHAIN so certificates contained in the message are not used as untrusted CAs")
	smimeCmd.Flags().BoolS("nodetach", "nodetach", false, "Use opaque signing")
	smimeCmd.Flags().BoolS("noindef", "noindef", false, "Disable CMS streaming")
	smimeCmd.Flags().BoolS("nointern", "nointern", false, "Don't search certificates in message for signer")
	smimeCmd.Flags().BoolS("nosigs", "nosigs", false, "Don't verify message signature")
	smimeCmd.Flags().BoolS("nosmimecap", "nosmimecap", false, "Omit the SMIMECapabilities attribute")
	smimeCmd.Flags().BoolS("noverify", "noverify", false, "Don't verify signers certificate")
	smimeCmd.Flags().StringS("out", "out", "", "Output file")
	smimeCmd.Flags().StringS("outform", "outform", "", "Output format SMIME (default), PEM or DER")
	smimeCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	smimeCmd.Flags().BoolS("pk7out", "pk7out", false, "Output PKCS#7 structure")
	smimeCmd.Flags().StringS("recip", "recip", "", "Recipient certificate file for decryption")
	smimeCmd.Flags().BoolS("resign", "resign", false, "Resign a signed message")
	smimeCmd.Flags().BoolS("sign", "sign", false, "Sign message")
	smimeCmd.Flags().StringS("signer", "signer", "", "Signer certificate file")
	smimeCmd.Flags().BoolS("stream", "stream", false, "Enable CMS streaming")
	smimeCmd.Flags().StringS("subject", "subject", "", "Subject")
	smimeCmd.Flags().BoolS("text", "text", false, "Include or delete text MIME headers")
	smimeCmd.Flags().StringS("to", "to", "", "To address")
	smimeCmd.Flags().BoolS("verify", "verify", false, "Verify signed message")
	common.AddProviderFlags(smimeCmd)
	common.AddRandomStateFlags(smimeCmd)
	common.AddValidationFlags(smimeCmd)
	rootCmd.AddCommand(smimeCmd)

	carapace.Gen(smimeCmd).FlagCompletion(carapace.ActionMap{
		"CAfile":   carapace.ActionFiles(),
		"CApath":   carapace.ActionDirectories(),
		"certfile": carapace.ActionFiles(),
		"config":   carapace.ActionFiles(),
		"content":  carapace.ActionFiles(),
		"in":       carapace.ActionFiles(),
		"inform":   carapace.ActionValues("DER", "PEM", "SMIME"),
		"inkey":    carapace.ActionFiles(),
		"keyform":  carapace.ActionValues("DER", "PEM", "P12"),
		"out":      carapace.ActionFiles(),
		"outform":  carapace.ActionValues("DER", "PEM", "SMIME"),
		"recip":    carapace.ActionFiles(),
		"signer":   carapace.ActionFiles(),
	})
}
