package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_account_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_account_registerCmd).Standalone()

	help_accountCmd.AddCommand(help_account_registerCmd)
}
