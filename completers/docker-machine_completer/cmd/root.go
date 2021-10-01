package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-machine",
	Short: "Create and manage machines running Docker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("bugsnag-api-token", false, "BugSnag API token for crash reporting [$MACHINE_BUGSNAG_API_TOKEN]")
	rootCmd.Flags().BoolP("debug", "D", false, "Enable debug mode")
	rootCmd.Flags().Bool("github-api-token", false, "Token to use for requests to the Github API [$MACHINE_GITHUB_API_TOKEN]")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().Bool("native-ssh", false, "Use the native (Go-based) SSH implementation. [$MACHINE_NATIVE_SSH]")
	rootCmd.Flags().StringP("storage-path", "s", "", "Configures storage path [$MACHINE_STORAGE_PATH]")
	rootCmd.Flags().Bool("tls-ca-cert", false, "CA to verify remotes against [$MACHINE_TLS_CA_CERT]")
	rootCmd.Flags().Bool("tls-ca-key", false, "Private key to generate certificates [$MACHINE_TLS_CA_KEY]")
	rootCmd.Flags().Bool("tls-client-cert", false, "Client cert to use for TLS [$MACHINE_TLS_CLIENT_CERT]")
	rootCmd.Flags().Bool("tls-client-key", false, "Private key used in client TLS auth [$MACHINE_TLS_CLIENT_KEY]")
	rootCmd.Flags().BoolP("version", "v", false, "print the version")
}
