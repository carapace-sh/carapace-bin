package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identity_provider_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List identity providers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identity_provider_listCmd).Standalone()

	identity_provider_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	identity_provider_listCmd.Flags().Bool("enabled", false, "List only enabled identity providers")
	identity_provider_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	identity_provider_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	identity_provider_listCmd.Flags().String("id", "", "Filter identity providers by ID")
	identity_provider_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	identity_provider_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	identity_provider_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	identity_provider_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	identity_provider_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	identity_provider_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	identity_provider_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	identity_providerCmd.AddCommand(identity_provider_listCmd)
}
