package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate the CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().String("auth-url", "", "Auth service URL with scheme (e.g. `ucs-auth://auth.example.com`). Required when logging in with `--token` outside a repository without a remote-url")
	loginCmd.Flags().BoolP("help", "h", false, "Print help")
	loginCmd.Flags().Bool("no-browser", false, "Avoid opening a browser to login")
	loginCmd.Flags().String("token", "", "Token value for non-interactive login (requires --token-type)")
	loginCmd.Flags().String("token-type", "", "Token type for non-interactive login (e.g. \"api-key\", \"eg1\", \"lore\")")
	rootCmd.AddCommand(loginCmd)

	carapace.Gen(loginCmd).FlagCompletion(carapace.ActionMap{
		"token-type": carapace.ActionValues("api-key", "eg1", "lore"),
	})

	carapace.Gen(loginCmd).PositionalCompletion(
		carapace.ActionValues(), // remote-url
	)
}
