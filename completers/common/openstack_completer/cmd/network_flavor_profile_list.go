package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_flavor_profile_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List network flavor profile(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_profile_listCmd).Standalone()

	network_flavor_profile_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_flavor_profile_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_flavor_profile_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_flavor_profile_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	network_flavor_profile_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	network_flavor_profile_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	network_flavor_profile_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_flavor_profile_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_flavor_profile_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_flavor_profile_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_flavor_profile_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_flavor_profile_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_flavor_profile_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_flavor_profileCmd.AddCommand(network_flavor_profile_listCmd)
}
