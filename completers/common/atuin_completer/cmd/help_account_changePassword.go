package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_account_changePasswordCmd = &cobra.Command{
	Use:   "change-password",
	Short: "Change your password",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_account_changePasswordCmd).Standalone()

	help_accountCmd.AddCommand(help_account_changePasswordCmd)
}
