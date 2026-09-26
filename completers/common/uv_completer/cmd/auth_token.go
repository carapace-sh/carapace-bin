package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var auth_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Show the authentication token for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_tokenCmd).Standalone()

	auth_tokenCmd.Flags().String("keyring-provider", "", "The keyring provider to use for reading credentials")
	auth_tokenCmd.Flags().StringP("username", "u", "", "The username to lookup")
	authCmd.AddCommand(auth_tokenCmd)
	carapace.Gen(auth_tokenCmd).FlagCompletion(carapace.ActionMap{
		"keyring-provider": uv.ActionKeyringProviders(),
	})
}
