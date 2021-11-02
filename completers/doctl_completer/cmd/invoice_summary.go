package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var invoice_summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Get a summary of an invoice",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoice_summaryCmd).Standalone()
	invoice_summaryCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ResourceID`, `ResourceUUID`, `Product`, `Description`, `GroupDescription`, `Amount`, `Duration`, `DurationUnit`, `StartTime`, `EndTime`, `ProjectName`, `Category`")
	invoice_summaryCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	invoiceCmd.AddCommand(invoice_summaryCmd)
}
