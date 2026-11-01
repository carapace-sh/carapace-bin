package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from your Anthropic account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_logoutCmd).Standalone()

	auth_logoutCmd.Flags().BoolP("help", "h", false, "Display help for command")
	authCmd.AddCommand(auth_logoutCmd)
}
