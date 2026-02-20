package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Store authentication information for a given host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_auth_loginCmd).Standalone()

	help_authCmd.AddCommand(help_auth_loginCmd)
}
