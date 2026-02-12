package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage your sync account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accountCmd).Standalone()

	accountCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(accountCmd)
}
