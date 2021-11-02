package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var invoice_pdfCmd = &cobra.Command{
	Use:   "pdf",
	Short: "Download a PDF file of an invoice",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(invoice_pdfCmd).Standalone()
	invoiceCmd.AddCommand(invoice_pdfCmd)
}
