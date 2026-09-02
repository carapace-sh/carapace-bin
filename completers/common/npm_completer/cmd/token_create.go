package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new authentication token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_createCmd).Standalone()
	token_createCmd.Flags().Bool("bypass-2fa", false, "bypass 2fa when creating a token")
	token_createCmd.Flags().StringArray("cidr", nil, "addresses used to limit access")
	token_createCmd.Flags().String("expires", "", "token expiration date")
	token_createCmd.Flags().String("name", "", "token name")
	token_createCmd.Flags().StringArray("orgs", nil, "orgs to grant access to")
	token_createCmd.Flags().String("orgs-permission", "", "permission level for orgs")
	token_createCmd.Flags().StringArray("packages", nil, "packages to grant access to")
	token_createCmd.Flags().Bool("packages-all", false, "grant access to all packages")
	token_createCmd.Flags().String("packages-and-scopes-permission", "", "permission level for packages and scopes")
	token_createCmd.Flags().String("password", "", "password for authentication")
	token_createCmd.Flags().Bool("read-only", false, "mark token as unable to publish")
	token_createCmd.Flags().StringArray("scopes", nil, "scopes to grant access to")
	token_createCmd.Flags().String("token-description", "", "token description")

	tokenCmd.AddCommand(token_createCmd)

	carapace.Gen(token_createCmd).FlagCompletion(carapace.ActionMap{
		"orgs-permission":                carapace.ActionValues("read-only", "read-write", "no-access"),
		"packages-and-scopes-permission": carapace.ActionValues("read-only", "read-write", "no-access"),
	})
}
