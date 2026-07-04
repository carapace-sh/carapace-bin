package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Print account information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accountCmd).Standalone()
	rootCmd.AddCommand(accountCmd)
}
