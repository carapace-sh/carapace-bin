package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_help_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_help_logoutCmd).Standalone()

	account_helpCmd.AddCommand(account_help_logoutCmd)
}
