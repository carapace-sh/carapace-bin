package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_loginCmd).Standalone()

	auth_loginCmd.Flags().String("keyring-provider", "", "The keyring provider to use for storage of credentials")
	auth_loginCmd.Flags().String("password", "", "The password to use for the service")
	auth_loginCmd.Flags().StringP("token", "t", "", "The token to use for the service")
	auth_loginCmd.Flags().StringP("username", "u", "", "The username to use for the service")
	authCmd.AddCommand(auth_loginCmd)
	carapace.Gen(auth_loginCmd).FlagCompletion(carapace.ActionMap{
		"keyring-provider": uv.ActionKeyringProviders(),
	})
}
