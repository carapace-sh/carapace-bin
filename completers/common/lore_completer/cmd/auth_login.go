package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate the CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_loginCmd).Standalone()

	auth_loginCmd.Flags().String("auth-url", "", "Auth service URL with scheme (e.g. `ucs-auth://auth.example.com`). Required when logging in with `--token` outside a repository without a remote-url")
	auth_loginCmd.Flags().BoolP("help", "h", false, "Print help")
	auth_loginCmd.Flags().Bool("no-browser", false, "Avoid opening a browser to login")
	auth_loginCmd.Flags().String("token", "", "Token value for non-interactive login (requires --token-type)")
	auth_loginCmd.Flags().String("token-type", "", "Token type for non-interactive login (e.g. \"api-key\", \"eg1\", \"lore\")")
	authCmd.AddCommand(auth_loginCmd)
}
