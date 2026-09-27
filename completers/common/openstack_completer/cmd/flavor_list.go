package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flavor_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List flavors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flavor_listCmd).Standalone()

	flavor_listCmd.Flags().Bool("all", false, "List all flavors, whether public or private")
	flavor_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	flavor_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	flavor_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	flavor_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	flavor_listCmd.Flags().Bool("long", false, "List additional fields in output")
	flavor_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	flavor_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	flavor_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	flavor_listCmd.Flags().String("min-disk", "", "Filters the flavors by a minimum disk space, in GiB.")
	flavor_listCmd.Flags().String("min-ram", "", "Filters the flavors by a minimum RAM, in MiB.")
	flavor_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	flavor_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	flavor_listCmd.Flags().Bool("private", false, "List only private flavors")
	flavor_listCmd.Flags().Bool("public", false, "List only public flavors (default)")
	flavor_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	flavor_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	flavor_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	flavor_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	flavorCmd.AddCommand(flavor_listCmd)
}
