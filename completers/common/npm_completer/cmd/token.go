package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage your authentication tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tokenCmd).Standalone()

	tokenCmd.PersistentFlags().Bool("bypass-2fa", false, "bypass two-factor authentication")
	tokenCmd.PersistentFlags().String("cidr", "", "CIDR whitelist for token")
	tokenCmd.PersistentFlags().String("expires", "", "token expiration in days")
	tokenCmd.PersistentFlags().String("name", "", "token name")
	tokenCmd.PersistentFlags().String("orgs", "", "orgs for token")
	tokenCmd.PersistentFlags().String("orgs-permission", "", "permission for orgs")
	tokenCmd.PersistentFlags().String("otp", "", "one-time password")
	tokenCmd.PersistentFlags().String("packages", "", "packages for token")
	tokenCmd.PersistentFlags().Bool("packages-all", false, "all packages for token")
	tokenCmd.PersistentFlags().String("packages-and-scopes-permission", "", "permission for packages and scopes")
	tokenCmd.PersistentFlags().String("password", "", "password for authentication")
	tokenCmd.PersistentFlags().Bool("read-only", false, "read-only token")
	tokenCmd.PersistentFlags().String("registry", "", "base URL of the npm registry")
	tokenCmd.PersistentFlags().String("scopes", "", "scopes for token")
	tokenCmd.PersistentFlags().String("token-description", "", "description for the token")
	rootCmd.AddCommand(tokenCmd)

	carapace.Gen(tokenCmd).FlagCompletion(carapace.ActionMap{
		"orgs-permission":                carapace.ActionValues("read-only", "read-write", "no-access"),
		"packages-and-scopes-permission": carapace.ActionValues("read-only", "read-write", "no-access"),
	})
}
