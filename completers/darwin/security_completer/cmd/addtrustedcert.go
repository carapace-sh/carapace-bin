package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var addTrustedCertCmd = &cobra.Command{
	Use:   "add-trusted-cert",
	Short: "Add trusted certificate(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addTrustedCertCmd).Standalone()
	addTrustedCertCmd.Flags().BoolP("admin", "d", false, "Add to admin cert store (default: user)")
	addTrustedCertCmd.Flags().StringP("allowed-error", "e", "", "Allowed error (integer, certExpired, hostnameMismatch)")
	addTrustedCertCmd.Flags().StringP("app-path", "a", "", "Specify application constraint")
	addTrustedCertCmd.Flags().BoolP("default", "D", false, "Add default setting instead of per-cert setting")
	addTrustedCertCmd.Flags().StringP("key-usage", "u", "", "Specify key usage (integer)")
	addTrustedCertCmd.Flags().StringP("keychain", "k", "", "Specify keychain to which cert is added")
	addTrustedCertCmd.Flags().StringP("policy", "p", "", "Policy constraint: ssl, smime, codeSign, etc.")
	addTrustedCertCmd.Flags().StringP("policy-string", "s", "", "Specify policy-specific string")
	addTrustedCertCmd.Flags().StringP("result-type", "r", "", "trustRoot, trustAsRoot, deny, or unspecified")
	addTrustedCertCmd.Flags().StringP("settings-in", "i", "", "Input trust settings file (default: user domain)")
	addTrustedCertCmd.Flags().StringP("settings-out", "o", "", "Output trust settings file (default: user domain)")
	rootCmd.AddCommand(addTrustedCertCmd)
	carapace.Gen(addTrustedCertCmd).PositionalCompletion(carapace.ActionFiles())
	carapace.Gen(addTrustedCertCmd).FlagCompletion(carapace.ActionMap{
		"keychain":     security.ActionKeychains(),
		"policy":       carapace.ActionValues("ssl", "smime", "codeSign", "IPSec", "iChat", "basic", "swUpdate", "pkgSign", "pkinitClient", "pkinitServer", "eap"),
		"result-type":  carapace.ActionValues("trustRoot", "trustAsRoot", "deny", "unspecified"),
		"settings-in":  carapace.ActionFiles(),
		"settings-out": carapace.ActionFiles(),
	})
}
