package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var verifyCertCmd = &cobra.Command{
	Use:   "verify-cert",
	Short: "Verify certificate(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyCertCmd).Standalone()
	verifyCertCmd.Flags().StringP("cert", "c", "", "Certificate to verify (DER or PEM)")
	verifyCertCmd.Flags().StringP("email", "e", "", "Specify email address for smime policy")
	verifyCertCmd.Flags().StringP("keychain", "k", "", "Keychain for intermediate certs")
	verifyCertCmd.Flags().BoolP("leaf-is-ca", "l", false, "Leaf certificate is a CA cert")
	verifyCertCmd.Flags().BoolP("local-only", "L", false, "Use local certificates only (no network fetch)")
	verifyCertCmd.Flags().BoolP("no-keychains", "n", false, "Avoid searching any keychains")
	verifyCertCmd.Flags().StringP("policy", "p", "", "Verification policy (default: basic)")
	verifyCertCmd.Flags().BoolP("quiet", "q", false, "Quiet, no stdout or stderr")
	verifyCertCmd.Flags().StringP("root-cert", "r", "", "Root certificate (DER or PEM)")
	verifyCertCmd.Flags().StringP("ssl-host", "s", "", "Specify SSL host name for ssl policy")
	rootCmd.AddCommand(verifyCertCmd)
	carapace.Gen(verifyCertCmd).PositionalCompletion(carapace.ActionFiles())
	carapace.Gen(verifyCertCmd).FlagCompletion(carapace.ActionMap{
		"cert":      carapace.ActionFiles(),
		"keychain":  security.ActionKeychains(),
		"policy":    carapace.ActionValues("ssl", "smime", "codeSign", "IPSec", "iChat", "basic", "swUpdate", "pkgSign", "pkinitClient", "pkinitServer", "eap", "appleID", "macappstore", "timestamping"),
		"root-cert": carapace.ActionFiles(),
	})
}
