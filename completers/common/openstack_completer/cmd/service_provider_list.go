package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_provider_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List service providers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_provider_listCmd).Standalone()

	service_provider_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	service_provider_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	service_provider_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	service_provider_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	service_provider_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	service_provider_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	service_provider_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	service_provider_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	service_provider_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	service_provider_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	service_providerCmd.AddCommand(service_provider_listCmd)
}
