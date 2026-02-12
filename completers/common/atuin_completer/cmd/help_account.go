package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage your sync account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_accountCmd).Standalone()

	helpCmd.AddCommand(help_accountCmd)
}
