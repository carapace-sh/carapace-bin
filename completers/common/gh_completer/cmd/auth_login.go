package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_loginCmd).Standalone()

	auth_loginCmd.Flags().BoolP("clipboard", "c", false, "Copy one-time OAuth device code to clipboard")
	auth_loginCmd.Flags().StringP("git-protocol", "p", "", "The protocol to use for git operations on this host: {ssh|https}")
	auth_loginCmd.Flags().StringP("hostname", "h", "", "The hostname of the GitHub instance to authenticate with")
	auth_loginCmd.Flags().Bool("insecure-storage", false, "Save authentication credentials in plain text instead of credential store")
	auth_loginCmd.Flags().StringSliceP("scopes", "s", nil, "Additional authentication scopes to request")
	auth_loginCmd.Flags().Bool("secure-storage", false, "Save authentication credentials in secure credential store")
	auth_loginCmd.Flags().Bool("skip-ssh-key", false, "Skip generate/upload SSH key prompt")
	auth_loginCmd.Flags().BoolP("web", "w", false, "Open a browser to authenticate")
	auth_loginCmd.Flags().Bool("with-token", false, "Read token from standard input")
	auth_loginCmd.Flag("secure-storage").Hidden = true
	authCmd.AddCommand(auth_loginCmd)

	carapace.Gen(auth_loginCmd).FlagCompletion(carapace.ActionMap{
		"git-protocol": carapace.ActionValues("ssh", "https"),
		"hostname":     action.ActionConfigHosts(),
		"scopes":       gh.ActionAuthScopes().MultiParts(":").UniqueList(","),
	})
}
