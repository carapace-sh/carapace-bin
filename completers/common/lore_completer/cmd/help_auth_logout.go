package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_auth_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored authentication and authorization tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_auth_logoutCmd).Standalone()

	help_authCmd.AddCommand(help_auth_logoutCmd)
}
