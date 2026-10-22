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
	auth_loginCmd.Flags().Bool("oauth", false, "Log in interactively with OAuth/OIDC (browser or device code)")
	auth_loginCmd.Flags().String("oauth-audience", "", "OAuth audience to request, so the access token is also accepted by that API (provider-specific)")
	auth_loginCmd.Flags().String("oauth-client-id", "", "OAuth client ID (defaults to \"rattler\")")
	auth_loginCmd.Flags().String("oauth-client-secret", "", "OAuth client secret (for confidential clients)")
	auth_loginCmd.Flags().String("oauth-flow", "", "OAuth flow: device-code (default), auth-code, auto")
	auth_loginCmd.Flags().String("oauth-issuer-url", "", "OIDC issuer URL (defaults to <https://{host>})")
	auth_loginCmd.Flags().String("oauth-redirect-uri", "", "OAuth redirect URI (defaults to a random localhost port). Set this when the OAuth client on the `IdP` side is registered with a specific redirect URI such as `http://127.0.0.1:8000/auth/oidc`")
	auth_loginCmd.Flags().StringSlice("oauth-scope", nil, "Additional OAuth scopes to request (repeatable)")
	auth_loginCmd.Flags().String("password", "", "The password to use (for basic HTTP authentication)")
	auth_loginCmd.Flags().String("s3-access-key-id", "", "The S3 access key ID")
	auth_loginCmd.Flags().String("s3-secret-access-key", "", "The S3 secret access key")
	auth_loginCmd.Flags().String("s3-session-token", "", "The S3 session token")
	auth_loginCmd.Flags().String("token", "", "The token to use (for authentication with prefix.dev)")
	auth_loginCmd.Flags().String("user-agent", "", "User-Agent header used for requests")
	auth_loginCmd.Flags().String("username", "", "The username to use (for basic HTTP authentication)")
	auth_loginCmd.Flags().Bool("workload-identity", false, "Log in non-interactively by exchanging the CI provider's OIDC ID token for an access token")
	auth_loginCmd.Flags().String("workload-identity-audience", "", "Audience requested in the OIDC ID token (defaults to the host, or prefix.dev for prefix.dev hosts)")
	auth_loginCmd.Flags().String("workload-identity-exchange", "", "How the server exchanges the ID token (defaults to prefix for prefix.dev hosts and token-exchange otherwise)")
	auth_loginCmd.Flags().String("workload-identity-token-endpoint", "", "Token endpoint for --workload-identity-exchange token-exchange: a URL or a path on the host")
	authCmd.AddCommand(auth_loginCmd)

	carapace.Gen(auth_loginCmd).FlagCompletion(carapace.ActionMap{
		"oauth-flow":                 carapace.ActionValues("device-code", "auth-code", "auto"),
		"workload-identity-exchange": carapace.ActionValues("token-exchange", "prefix"),
	})
}
