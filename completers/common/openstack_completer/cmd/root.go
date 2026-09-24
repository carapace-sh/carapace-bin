package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openstack",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("debug", false, "Show tracebacks on errors.")
	rootCmd.Flags().Bool("insecure", false, "Disable server certificate verification")
	rootCmd.Flags().String("log-file", "", "Specify a file to log output.")
	rootCmd.Flags().String("os-access-token", "", "With v3oidcaccesstoken: OAuth 2.0 Access Token")
	rootCmd.Flags().String("os-access-token-endpoint", "", "With v3oidcdeviceauthz: OpenID Connect Provider Token Endpoint.")
	rootCmd.Flags().String("os-access-token-type", "", "With v3oidcdeviceauthz: OAuth 2.0 Authorization Server Introspection token type, it is used to decide which type of token will be used when processing token introspection.")
	rootCmd.Flags().String("os-application-credential-id", "", "With v3applicationcredential: Application credential ID")
	rootCmd.Flags().String("os-application-credential-name", "", "With v3applicationcredential: Application credential name")
	rootCmd.Flags().String("os-application-credential-secret", "", "With v3applicationcredential: Application credential auth secret")
	rootCmd.Flags().String("os-auth-methods", "", "With v3multifactor: Methods to authenticate with.")
	rootCmd.Flags().String("os-auth-type", "", "Select an authentication type.")
	rootCmd.Flags().String("os-auth-url", "", "With v3applicationcredential: Authentication URL")
	rootCmd.Flags().Bool("os-beta-command", false, "Enable beta commands which are subject to change")
	rootCmd.Flags().String("os-cacert", "", "CA certificate bundle file (Env: OS_CACERT)")
	rootCmd.Flags().String("os-cert", "", "Client certificate bundle file (Env: OS_CERT)")
	rootCmd.Flags().String("os-client-id", "", "With v3oidcdeviceauthz: OAuth 2.0 Client ID")
	rootCmd.Flags().String("os-client-secret", "", "With v3oidcdeviceauthz: OAuth 2.0 Client Secret")
	rootCmd.Flags().String("os-cloud", "", "Cloud name in clouds.yaml (Env: OS_CLOUD)")
	rootCmd.Flags().String("os-code", "", "With v3oidcauthcode: OAuth 2.0 Authorization Code")
	rootCmd.Flags().String("os-code-challenge-method", "", "With v3oidcdeviceauthz: PKCE Challenge Method (RFC 7636)")
	rootCmd.Flags().String("os-compute-api-version", "", "Compute API version, default=2.1 (Env: OS_COMPUTE_API_VERSION)")
	rootCmd.Flags().String("os-default-domain", "", "Default domain ID, default=default.")
	rootCmd.Flags().String("os-default-domain-id", "", "With password: Optional domain ID to use with v3 and v2 parameters.")
	rootCmd.Flags().String("os-default-domain-name", "", "With password: Optional domain name to use with v3 API and v2 parameters.")
	rootCmd.Flags().String("os-device-authorization-endpoint", "", "With v3oidcdeviceauthz: OAuth 2.0 Device Authorization Endpoint.")
	rootCmd.Flags().String("os-discovery-endpoint", "", "With v3oidcdeviceauthz: OpenID Connect Discovery Document URL.")
	rootCmd.Flags().String("os-domain-id", "", "With v3applicationcredential: Domain ID to scope to")
	rootCmd.Flags().String("os-domain-name", "", "With v3applicationcredential: Domain name to scope to")
	rootCmd.Flags().String("os-endpoint", "", "With admin_token: The endpoint that will always be used")
	rootCmd.Flags().String("os-endpoint-override", "", "Use this API endpoint instead of the Service Catalog.")
	rootCmd.Flags().String("os-identity-api-version", "", "Identity API version, default=3 (Env: OS_IDENTITY_API_VERSION)")
	rootCmd.Flags().String("os-identity-provider", "", "With v3oidcaccesstoken: Identity Provider's name")
	rootCmd.Flags().String("os-idp-otp-key", "", "With v3oidcpassword: A key to be used in the Identity Provider access token endpoint to pass the OTP value.")
	rootCmd.Flags().String("os-image-api-version", "", "Image API version, default=2 (Env: OS_IMAGE_API_VERSION)")
	rootCmd.Flags().String("os-interface", "", "Select an interface type.")
	rootCmd.Flags().String("os-key", "", "Client certificate key file (Env: OS_KEY)")
	rootCmd.Flags().String("os-network-api-version", "", "Network API version, default=2.0 (Env: OS_NETWORK_API_VERSION)")
	rootCmd.Flags().String("os-oauth2-client-id", "", "With v3oauth2clientcredential: Client id for OAuth2.0")
	rootCmd.Flags().String("os-oauth2-client-secret", "", "With v3oauth2clientcredential: Client secret for OAuth2.0")
	rootCmd.Flags().String("os-oauth2-endpoint", "", "With v3oauth2clientcredential: Endpoint for OAuth2.0")
	rootCmd.Flags().String("os-object-api-version", "", "Object API version, default=1 (Env: OS_OBJECT_API_VERSION)")
	rootCmd.Flags().String("os-openid-scope", "", "With v3oidcdeviceauthz: OpenID Connect scope that is requested from authorization server.")
	rootCmd.Flags().String("os-passcode", "", "With v3totp: User's TOTP passcode")
	rootCmd.Flags().String("os-password", "", "With password: User's password")
	rootCmd.Flags().String("os-project-domain-id", "", "With v3applicationcredential: Domain ID containing project")
	rootCmd.Flags().String("os-project-domain-name", "", "With v3applicationcredential: Domain name containing project")
	rootCmd.Flags().String("os-project-id", "", "With v3applicationcredential: Project ID to scope to")
	rootCmd.Flags().String("os-project-name", "", "With v3applicationcredential: Project name to scope to")
	rootCmd.Flags().String("os-protocol", "", "With v3oidcaccesstoken: Protocol for federated plugin")
	rootCmd.Flags().String("os-redirect-uri", "", "With v3oidcauthcode: OpenID Connect Redirect URL")
	rootCmd.Flags().String("os-region-name", "", "Authentication region name (Env: OS_REGION_NAME)")
	rootCmd.Flags().String("os-remote-project-domain-id", "", "Domain ID of the project when authenticating to a service provider if using Keystone-to-Keystone federation.")
	rootCmd.Flags().String("os-remote-project-domain-name", "", "Domain name of the project when authenticating to a service provider if using Keystone-to-Keystone federation.")
	rootCmd.Flags().String("os-remote-project-id", "", "Project ID when authenticating to a service provider if using Keystone-to-Keystone federation.")
	rootCmd.Flags().String("os-remote-project-name", "", "Project name when authenticating to a service provider if using Keystone-to-Keystone federation.")
	rootCmd.Flags().String("os-service-provider", "", "Authenticate with and perform the command on a service provider using Keystone-to-keystone federation.")
	rootCmd.Flags().String("os-share-api-version", "", "Shared File System API version, default=2.96 (version supported by both the client and the server) (Env: OS_SHARE_API_VERSION)")
	rootCmd.Flags().String("os-system-scope", "", "With v3applicationcredential: Scope for system operations")
	rootCmd.Flags().String("os-tenant-id", "", "==SUPPRESS==")
	rootCmd.Flags().String("os-tenant-name", "", "==SUPPRESS==")
	rootCmd.Flags().String("os-token", "", "With token: Token to authenticate with")
	rootCmd.Flags().String("os-trust-id", "", "With v3applicationcredential: ID of the trust to use as a trustee use")
	rootCmd.Flags().String("os-user-domain-id", "", "With v3applicationcredential: User's domain ID")
	rootCmd.Flags().String("os-user-domain-name", "", "With v3applicationcredential: User's domain name")
	rootCmd.Flags().String("os-user-id", "", "With v3applicationcredential: User's user ID")
	rootCmd.Flags().String("os-username", "", "With v3applicationcredential: User's username")
	rootCmd.Flags().String("os-volume-api-version", "", "Volume API version, default=3 (Env: OS_VOLUME_API_VERSION)")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress output except warnings and errors.")
	rootCmd.Flags().Bool("timing", false, "Print API call timing info")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase verbosity of output.")
	rootCmd.Flags().Bool("verify", false, "Verify server certificate (default)")
	rootCmd.Flags().String("version", "", "show program's version number and exit")
}
