package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var cmsCmd = &cobra.Command{
	Use:   "cms",
	Short: "Create or decode CMS messages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cmsCmd).Standalone()
	cmsCmd.Flags().StringP("cert-usage", "u", "", "Set type of certificate usage (0-11)")
	cmsCmd.Flags().StringP("content", "c", "", "Use this detached content file (decode)")
	cmsCmd.Flags().BoolP("create-encrypted", "C", false, "Create a CMS encrypted message")
	cmsCmd.Flags().BoolP("create-enveloped", "E", false, "Create a CMS enveloped message")
	cmsCmd.Flags().BoolP("create-signed", "S", false, "Create a CMS signed message")
	cmsCmd.Flags().BoolP("decode", "D", false, "Decode a CMS message")
	cmsCmd.Flags().StringP("enc-key-pref", "Y", "", "Include an EncryptionKeyPreference attribute (encode)")
	cmsCmd.Flags().StringP("envelope", "e", "", "Specify envelope file")
	cmsCmd.Flags().StringP("hash-algorithm", "H", "", "Hash algorithm (encode, default: SHA1)")
	cmsCmd.Flags().StringP("headers", "h", "", "Generate email headers with info (decode)")
	cmsCmd.Flags().StringP("input", "i", "", "Use infile as source of data (default: stdin)")
	cmsCmd.Flags().StringP("keychain", "k", "", "Specify keychain to use")
	cmsCmd.Flags().BoolP("no-content", "T", false, "Do not include content in CMS message (encode)")
	cmsCmd.Flags().StringP("output", "o", "", "Use outfile as destination of data (default: stdout)")
	cmsCmd.Flags().StringP("password", "p", "", "Use password as key db password")
	cmsCmd.Flags().StringP("recipients", "r", "", "Create envelope for recipients (encode)")
	cmsCmd.Flags().StringP("signer", "N", "", "Use certificate named for signing (encode)")
	cmsCmd.Flags().BoolP("signing-time", "G", false, "Include a signing time attribute (encode)")
	cmsCmd.Flags().BoolP("single-byte", "s", false, "Pass data a single byte at a time to CMS")
	cmsCmd.Flags().BoolP("smime-capabilities", "P", false, "Include SMIMECapabilities attribute (encode)")
	cmsCmd.Flags().StringP("subject-key-id", "Z", "", "Find a certificate by subject key ID (encode)")
	cmsCmd.Flags().BoolP("suppress-content", "n", false, "Suppress output of content (decode)")
	cmsCmd.Flags().BoolP("verbose", "v", false, "Print debugging information")
	rootCmd.AddCommand(cmsCmd)
	carapace.Gen(cmsCmd).FlagCompletion(carapace.ActionMap{
		"content":        carapace.ActionFiles(),
		"envelope":       carapace.ActionFiles(),
		"hash-algorithm": carapace.ActionValues("MD2", "MD4", "MD5", "SHA1", "SHA256", "SHA384", "SHA512"),
		"input":          carapace.ActionFiles(),
		"keychain":       security.ActionKeychains(),
		"output":         carapace.ActionFiles(),
	})
}
