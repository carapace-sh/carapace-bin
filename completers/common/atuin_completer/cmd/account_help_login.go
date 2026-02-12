package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_help_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the configured server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_help_loginCmd).Standalone()

	account_helpCmd.AddCommand(account_help_loginCmd)
}
