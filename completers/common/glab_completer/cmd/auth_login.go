package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with a GitLab instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_loginCmd).Standalone()

	auth_loginCmd.Flags().StringP("api-host", "a", "", "Hostname for the API endpoint, if different from --hostname. Accepts a hostname or hostname:port. Use only when the API is served from a different host than the Git remote.")
	auth_loginCmd.Flags().StringP("api-protocol", "p", "", "API protocol. Options: https, http.")
	auth_loginCmd.Flags().String("container-registry-domains", "", "Container registry and image dependency proxy domains, comma-separated.")
	auth_loginCmd.Flags().Bool("device", false, "Use the OAuth 2.0 device authorization flow. Useful for headless environments where a local browser is not available. Requires GitLab 17.9 or later.")
	auth_loginCmd.Flags().StringP("git-protocol", "g", "", "Git protocol. Options: ssh, https, http.")
	auth_loginCmd.Flags().String("hostname", "", "The hostname of the GitLab instance to authenticate with.")
	auth_loginCmd.Flags().Bool("insecure-storage", false, "Store the token as plaintext in the configuration file instead of the operating system's keyring.")
	auth_loginCmd.Flags().StringP("job-token", "j", "", "CI job token.")
	auth_loginCmd.Flags().String("ssh-hostname", "", "SSH hostname for instances with a different SSH endpoint. A port is not required; Git uses the port from the remote URL.")
	auth_loginCmd.Flags().Bool("stdin", false, "Read the token from standard input.")
	auth_loginCmd.Flags().StringP("token", "t", "", "Your GitLab access token.")
	auth_loginCmd.Flags().Bool("use-keyring", false, "Store the token in your operating system's keyring.")
	auth_loginCmd.Flags().Bool("web", false, "Skip the login type prompt and use web/OAuth login.")
	auth_loginCmd.Flag("use-keyring").Hidden = true
	authCmd.AddCommand(auth_loginCmd)

	carapace.Gen(auth_loginCmd).FlagCompletion(carapace.ActionMap{
		"api-protocol": carapace.ActionValues("https", "http"),
		"git-protocol": carapace.ActionValues("ssh", "https", "http"),
		"hostname":     action.ActionConfigHosts(),
	})
}
