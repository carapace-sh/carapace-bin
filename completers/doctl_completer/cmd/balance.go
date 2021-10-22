package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Display commands for retrieving your account balance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(balanceCmd).Standalone()
	rootCmd.AddCommand(balanceCmd)
}
