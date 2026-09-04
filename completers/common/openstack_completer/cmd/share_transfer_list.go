package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_transfer_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all transfers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_transfer_listCmd).Standalone()

	share_transfer_listCmd.Flags().Bool("all-projects", false, "Shows details for all tenants.")
	share_transfer_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_transfer_listCmd.Flags().String("detailed", "", "Show detailed information about filtered share transfers.")
	share_transfer_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_transfer_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_transfer_listCmd.Flags().String("id", "", "Filter share transfers by ID.")
	share_transfer_listCmd.Flags().String("limit", "", "Maximum number of transfer records to return.")
	share_transfer_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_transfer_listCmd.Flags().String("name", "", "Filter share transfers by name.")
	share_transfer_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_transfer_listCmd.Flags().String("offset", "", "Start position of transfer records listing.")
	share_transfer_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_transfer_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_transfer_listCmd.Flags().String("resource-id", "", "Filter share transfers by resource ID.")
	share_transfer_listCmd.Flags().String("resource-type", "", "Filter share transfers by resource type, which can be share.")
	share_transfer_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_transfer_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_transfer_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_transfer_listCmd.Flags().String("sort-dir", "", "Sort direction, available values are ('asc', 'desc').")
	share_transfer_listCmd.Flags().String("sort-key", "", "Key to be sorted, available keys are ('id', 'resource_type', 'resource_id', 'name', 'source_project_id', 'destination_project_id', 'created_at', 'expires_at').")
	share_transfer_listCmd.Flags().String("source-project-id", "", "Filter share transfers by ID of the Project that initiated the transfer.")
	share_transferCmd.AddCommand(share_transfer_listCmd)
}
