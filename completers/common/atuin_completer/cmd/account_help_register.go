package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_help_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_help_registerCmd).Standalone()

	account_helpCmd.AddCommand(account_help_registerCmd)
}
