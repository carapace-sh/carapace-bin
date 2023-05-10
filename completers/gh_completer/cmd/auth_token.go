package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Print the auth token gh is configured to use",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_tokenCmd).Standalone()

	auth_tokenCmd.Flags().StringP("hostname", "h", "", "The hostname of the GitHub instance authenticated with")
	authCmd.AddCommand(auth_tokenCmd)

	carapace.Gen(auth_tokenCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
	})
}
