package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var user_password_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Change current user password",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_password_setCmd).Standalone()

	user_password_setCmd.Flags().String("original-password", "", "Original user password")
	user_password_setCmd.Flags().String("password", "", "New user password")
	user_passwordCmd.AddCommand(user_password_setCmd)
}
