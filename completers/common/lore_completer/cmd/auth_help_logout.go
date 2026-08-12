package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_help_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored authentication and authorization tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_help_logoutCmd).Standalone()

	auth_helpCmd.AddCommand(auth_help_logoutCmd)
}
