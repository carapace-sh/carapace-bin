package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_help_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate the CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_help_loginCmd).Standalone()

	auth_helpCmd.AddCommand(auth_help_loginCmd)
}
