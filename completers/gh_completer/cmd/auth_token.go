package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var auth_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Print the authentication token gh uses for a hostname and account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_tokenCmd).Standalone()

	auth_tokenCmd.Flags().StringP("hostname", "h", "", "The hostname of the GitHub instance authenticated with")
	auth_tokenCmd.Flags().Bool("secure-storage", false, "Search only secure credential store for authentication token")
	auth_tokenCmd.Flags().StringP("user", "u", "", "The account to log out of")
	auth_tokenCmd.Flag("secure-storage").Hidden = true
	authCmd.AddCommand(auth_tokenCmd)

	carapace.Gen(auth_tokenCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
		"user": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionConfigUsers(auth_tokenCmd.Flag("hostname").Value.String())
		}),
	})
}
