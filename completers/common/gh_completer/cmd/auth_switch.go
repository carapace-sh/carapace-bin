package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var auth_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch active GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_switchCmd).Standalone()

	auth_switchCmd.Flags().StringP("hostname", "h", "", "The hostname of the GitHub instance to switch account for")
	auth_switchCmd.Flags().StringP("user", "u", "", "The account to switch to")
	authCmd.AddCommand(auth_switchCmd)

	carapace.Gen(auth_switchCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
		"user": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionConfigUsers(auth_switchCmd.Flag("hostname").Value.String())
		}),
	})
}
