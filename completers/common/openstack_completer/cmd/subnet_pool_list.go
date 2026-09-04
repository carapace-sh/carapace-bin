package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subnet_pool_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List subnet pools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_pool_listCmd).Standalone()

	subnet_pool_listCmd.Flags().String("address-scope", "", "List only subnet pools with the specified address scope (name or ID)")
	subnet_pool_listCmd.Flags().String("any-tags", "", "List subnet pools which have any given tag(s) (Comma-separated list of tags)")
	subnet_pool_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	subnet_pool_listCmd.Flags().Bool("default", false, "List only subnet pools used as the default external subnet pool")
	subnet_pool_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	subnet_pool_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	subnet_pool_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	subnet_pool_listCmd.Flags().Bool("long", false, "List additional fields in output")
	subnet_pool_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	subnet_pool_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	subnet_pool_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	subnet_pool_listCmd.Flags().String("name", "", "List only subnet pools with the specified name")
	subnet_pool_listCmd.Flags().Bool("no-default", false, "List only subnet pools not used as the default external subnet pool")
	subnet_pool_listCmd.Flags().Bool("no-share", false, "List only subnet pools not shared between projects")
	subnet_pool_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	subnet_pool_listCmd.Flags().String("not-any-tags", "", "Exclude subnet pools which have any given tag(s) (Comma-separated list of tags)")
	subnet_pool_listCmd.Flags().String("not-tags", "", "Exclude subnet pools which have all given tag(s) (Comma-separated list of tags)")
	subnet_pool_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	subnet_pool_listCmd.Flags().String("project", "", "List only subnet pools with the specified project (name or ID)")
	subnet_pool_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	subnet_pool_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	subnet_pool_listCmd.Flags().Bool("share", false, "List only subnet pools shared between projects")
	subnet_pool_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	subnet_pool_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	subnet_pool_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	subnet_pool_listCmd.Flags().String("tags", "", "List subnet pools which have all given tag(s) (Comma-separated list of tags)")
	subnet_poolCmd.AddCommand(subnet_pool_listCmd)
}
