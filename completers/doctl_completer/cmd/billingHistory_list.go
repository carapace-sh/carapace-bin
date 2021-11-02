package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var billingHistory_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a paginated billing history for a user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingHistory_listCmd).Standalone()
	billingHistory_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Date`, `Type`, `Description`, `Amount`, `InvoiceID`, `InvoiceUUID`")
	billingHistory_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	billingHistoryCmd.AddCommand(billingHistory_listCmd)

	carapace.Gen(billingHistory_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("Date", "Type", "Description", "Amount", "InvoiceID", "InvoiceUUID").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
