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

	auth_loginCmd.Flags().StringP("api-host", "a", "", "API host url.")
	auth_loginCmd.Flags().StringP("api-protocol", "p", "", "API protocol: https, http")
	auth_loginCmd.Flags().StringP("git-protocol", "g", "", "Git protocol: ssh, https, http")
	auth_loginCmd.Flags().StringP("hostname", "h", "", "The hostname of the GitLab instance to authenticate with.")
	auth_loginCmd.Flags().StringP("job-token", "j", "", "CI job token.")
	auth_loginCmd.Flags().Bool("stdin", false, "Read token from standard input.")
	auth_loginCmd.Flags().StringP("token", "t", "", "Your GitLab access token.")
	auth_loginCmd.Flags().Bool("use-keyring", false, "Store token in your operating system's keyring.")
	authCmd.AddCommand(auth_loginCmd)

	carapace.Gen(auth_loginCmd).FlagCompletion(carapace.ActionMap{
		"api-protocol": carapace.ActionValues("https", "http"),
		"git-protocol": carapace.ActionValues("ssh", "https", "http"),
		"hostname":     action.ActionConfigHosts(),
	})
}
