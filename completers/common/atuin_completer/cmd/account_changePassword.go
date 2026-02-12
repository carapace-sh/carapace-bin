package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_changePasswordCmd = &cobra.Command{
	Use:   "change-password",
	Short: "Change your password",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_changePasswordCmd).Standalone()

	account_changePasswordCmd.Flags().StringP("current-password", "c", "", "")
	account_changePasswordCmd.Flags().BoolP("help", "h", false, "Print help")
	account_changePasswordCmd.Flags().StringP("new-password", "n", "", "")
	accountCmd.AddCommand(account_changePasswordCmd)
}
