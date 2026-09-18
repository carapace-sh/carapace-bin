package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_service_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List share services (Admin only).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_service_listCmd).Standalone()

	share_service_listCmd.Flags().String("binary", "", "Filter services by the name of the service.")
	share_service_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_service_listCmd.Flags().String("ensuring", "", "Filter services running ensure shares or not.")
	share_service_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_service_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_service_listCmd.Flags().String("host", "", "Filter services by name of the host.")
	share_service_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_service_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_service_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_service_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_service_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_service_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_service_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_service_listCmd.Flags().String("state", "", "Filter results by state.")
	share_service_listCmd.Flags().String("status", "", "Filter results by status.")
	share_service_listCmd.Flags().String("zone", "", "Filter services by their availability zone.")
	share_serviceCmd.AddCommand(share_service_listCmd)
}
