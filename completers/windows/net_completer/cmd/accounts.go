package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "update user accounts and password policies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accountsCmd).Standalone()
	rootCmd.AddCommand(accountsCmd)
}
