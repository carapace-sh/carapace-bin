package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate the CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_loginCmd).Standalone()

	helpCmd.AddCommand(help_loginCmd)
}
