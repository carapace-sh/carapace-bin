package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Store authentication information for a given host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_loginCmd).Standalone()

	auth_loginCmd.Flags().String("conda-token", "", "The token to use on anaconda.org / quetz authentication")
	auth_loginCmd.Flags().Bool("oauth", false, "Use OAuth/OIDC authentication")
	auth_loginCmd.Flags().String("oauth-client-id", "", "OAuth client ID (defaults to \"rattler\")")
	auth_loginCmd.Flags().String("oauth-client-secret", "", "OAuth client secret (for confidential clients)")
	auth_loginCmd.Flags().String("oauth-flow", "", "OAuth flow: device-code (default), auth-code, auto")
	auth_loginCmd.Flags().String("oauth-issuer-url", "", "OIDC issuer URL (defaults to <https://{host>)")
	auth_loginCmd.Flags().String("oauth-redirect-uri", "", "OAuth redirect URI (defaults to a random localhost port). Set this when the OAuth client on the `IdP` side is registered with a specific redirect URI such as `http://127.0.0.1:8000/auth/oidc`")
	auth_loginCmd.Flags().StringSlice("oauth-scope", nil, "Additional OAuth scopes to request (repeatable)")
	auth_loginCmd.Flags().String("password", "", "The password to use (for basic HTTP authentication)")
	auth_loginCmd.Flags().String("s3-access-key-id", "", "The S3 access key ID")
	auth_loginCmd.Flags().String("s3-secret-access-key", "", "The S3 secret access key")
	auth_loginCmd.Flags().String("s3-session-token", "", "The S3 session token")
	auth_loginCmd.Flags().String("token", "", "The token to use (for authentication with prefix.dev)")
	auth_loginCmd.Flags().String("user-agent", "", "User-Agent header used for requests")
	auth_loginCmd.Flags().String("username", "", "The username to use (for basic HTTP authentication)")
	authCmd.AddCommand(auth_loginCmd)

	carapace.Gen(auth_loginCmd).FlagCompletion(carapace.ActionMap{
		"oauth-flow": carapace.ActionValues("device-code", "auth-code", "auto"),
	})
}
