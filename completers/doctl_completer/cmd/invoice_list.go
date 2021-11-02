package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var invoice_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of the invoices for your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoice_listCmd).Standalone()
	invoice_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ResourceID`, `ResourceUUID`, `Product`, `Description`, `GroupDescription`, `Amount`, `Duration`, `DurationUnit`, `StartTime`, `EndTime`, `ProjectName`, `Category`")
	invoice_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	invoiceCmd.AddCommand(invoice_listCmd)
}
