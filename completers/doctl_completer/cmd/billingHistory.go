package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var billingHistoryCmd = &cobra.Command{
	Use:   "billing-history",
	Short: "Display commands for retrieving your billing history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingHistoryCmd).Standalone()
	rootCmd.AddCommand(billingHistoryCmd)
}
