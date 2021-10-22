package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var invoice_csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "Download a CSV file of an invoice",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoice_csvCmd).Standalone()
	invoiceCmd.AddCommand(invoice_csvCmd)
}
