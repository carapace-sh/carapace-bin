package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_help_changePasswordCmd = &cobra.Command{
	Use:   "change-password",
	Short: "Change your password",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_help_changePasswordCmd).Standalone()

	account_helpCmd.AddCommand(account_help_changePasswordCmd)
}
