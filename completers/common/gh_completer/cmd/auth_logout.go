package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var auth_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of a GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_logoutCmd).Standalone()

	auth_logoutCmd.Flags().StringP("hostname", "h", "", "The hostname of the GitHub instance to log out of")
	auth_logoutCmd.Flags().StringP("user", "u", "", "The account to log out of")
	authCmd.AddCommand(auth_logoutCmd)

	carapace.Gen(auth_logoutCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
		"user": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionConfigUsers(auth_logoutCmd.Flag("hostname").Value.String())
		}),
	})
}
