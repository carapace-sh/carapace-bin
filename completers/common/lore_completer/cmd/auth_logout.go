package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored authentication and authorization tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_logoutCmd).Standalone()

	auth_logoutCmd.Flags().String("auth-url", "", "Auth service URL (omit to use current repository's auth URL)")
	auth_logoutCmd.Flags().BoolP("help", "h", false, "Print help")
	auth_logoutCmd.Flags().String("resource", "", "Resource ID to remove a specific authorization (e.g. \"urc-{id}\")")
	auth_logoutCmd.Flags().String("user-id", "", "User ID to remove (omit to remove all identities)")
	authCmd.AddCommand(auth_logoutCmd)
}
