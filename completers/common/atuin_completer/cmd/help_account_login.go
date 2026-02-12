package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_account_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the configured server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_account_loginCmd).Standalone()

	help_accountCmd.AddCommand(help_account_loginCmd)
}
