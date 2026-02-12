package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_logoutCmd).Standalone()

	account_logoutCmd.Flags().BoolP("help", "h", false, "Print help")
	accountCmd.AddCommand(account_logoutCmd)
}
