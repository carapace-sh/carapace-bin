package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var useraccountCmd = &cobra.Command{
	Use:   "useraccount",
	Short: "user account management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(useraccountCmd).Standalone()
	rootCmd.AddCommand(useraccountCmd)
}
