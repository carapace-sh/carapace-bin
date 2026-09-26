package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var auth_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_logoutCmd).Standalone()

	auth_logoutCmd.Flags().String("keyring-provider", "", "The keyring provider to use for storage of credentials")
	auth_logoutCmd.Flags().StringP("username", "u", "", "The username to logout")
	authCmd.AddCommand(auth_logoutCmd)
	carapace.Gen(auth_logoutCmd).FlagCompletion(carapace.ActionMap{
		"keyring-provider": uv.ActionKeyringProviders(),
	})
}
