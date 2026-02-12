package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_registerCmd).Standalone()

	account_registerCmd.Flags().StringP("email", "e", "", "")
	account_registerCmd.Flags().BoolP("help", "h", false, "Print help")
	account_registerCmd.Flags().StringP("password", "p", "", "")
	account_registerCmd.Flags().StringP("username", "u", "", "")
	accountCmd.AddCommand(account_registerCmd)
}
