package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var invoiceCmd = &cobra.Command{
	Use:   "invoice",
	Short: "Display commands for retrieving invoices for your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoiceCmd).Standalone()
	rootCmd.AddCommand(invoiceCmd)
}
