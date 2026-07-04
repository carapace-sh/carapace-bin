package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findIdentityCmd = &cobra.Command{
	Use:   "find-identity",
	Short: "Find an identity (certificate + private key)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findIdentityCmd).Standalone()
	findIdentityCmd.Flags().StringP("policy", "p", "", "Specify policy (basic, ssl-client, ssl-server, smime, etc.)")
	findIdentityCmd.Flags().StringP("string", "s", "", "Specify optional policy-specific string")
	findIdentityCmd.Flags().BoolP("valid-only", "v", false, "Show valid identities only (default: show all)")
	rootCmd.AddCommand(findIdentityCmd)
	carapace.Gen(findIdentityCmd).FlagCompletion(carapace.ActionMap{
		"policy": carapace.ActionValues(
			"basic",
			"ssl-client",
			"ssl-server",
			"smime",
			"eap",
			"ipsec",
			"ichat",
			"codesigning",
			"sys-default",
			"sys-kerberos-kdc",
		),
	})
}
