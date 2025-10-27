package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from a GitLab instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_logoutCmd).Standalone()

	auth_logoutCmd.Flags().String("hostname", "", "The hostname of the GitLab instance.")
	auth_logoutCmd.MarkFlagRequired("hostname")
	authCmd.AddCommand(auth_logoutCmd)

	carapace.Gen(auth_logoutCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
	})
}
