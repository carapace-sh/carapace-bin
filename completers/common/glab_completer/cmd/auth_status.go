package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "View authentication status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_statusCmd).Standalone()

	auth_statusCmd.Flags().BoolP("all", "a", false, "Check the authentication status of all configured instances.")
	auth_statusCmd.Flags().String("hostname", "", "Check the authentication status of a specific instance.")
	auth_statusCmd.Flags().BoolP("show-token", "t", false, "Display the authentication token.")
	authCmd.AddCommand(auth_statusCmd)

	carapace.Gen(auth_statusCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
	})
}
