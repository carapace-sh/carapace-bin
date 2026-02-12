package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the configured server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_loginCmd).Standalone()

	account_loginCmd.Flags().BoolP("help", "h", false, "Print help")
	account_loginCmd.Flags().StringP("key", "k", "", "The encryption key for your account")
	account_loginCmd.Flags().StringP("password", "p", "", "")
	account_loginCmd.Flags().StringP("username", "u", "", "")
	accountCmd.AddCommand(account_loginCmd)
}
