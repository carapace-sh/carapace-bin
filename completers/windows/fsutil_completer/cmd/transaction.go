package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "transaction management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transactionCmd).Standalone()
	rootCmd.AddCommand(transactionCmd)
}
