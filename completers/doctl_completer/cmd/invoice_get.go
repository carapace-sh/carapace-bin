package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var invoice_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a list of all the items on an invoice",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoice_getCmd).Standalone()
	invoice_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ResourceID`, `ResourceUUID`, `Product`, `Description`, `GroupDescription`, `Amount`, `Duration`, `DurationUnit`, `StartTime`, `EndTime`, `ProjectName`, `Category`")
	invoice_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	invoiceCmd.AddCommand(invoice_getCmd)
}
