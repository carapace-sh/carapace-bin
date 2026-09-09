package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List routers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_listCmd).Standalone()

	router_listCmd.Flags().String("agent", "", "List only routers hosted by the specified agent (ID only)")
	router_listCmd.Flags().String("any-tags", "", "List routers which have any given tag(s) (Comma-separated list of tags)")
	router_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	router_listCmd.Flags().Bool("disable", false, "List disabled routers")
	router_listCmd.Flags().Bool("enable", false, "List enabled routers")
	router_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	router_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	router_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	router_listCmd.Flags().Bool("long", false, "List additional fields in output")
	router_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	router_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	router_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	router_listCmd.Flags().String("name", "", "List routers according to their name")
	router_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	router_listCmd.Flags().String("not-any-tags", "", "Exclude routers which have any given tag(s) (Comma-separated list of tags)")
	router_listCmd.Flags().String("not-tags", "", "Exclude routers which have all given tag(s) (Comma-separated list of tags)")
	router_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	router_listCmd.Flags().String("project", "", "List only routers with the specified project (name or ID)")
	router_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	router_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	router_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	router_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	router_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	router_listCmd.Flags().String("tags", "", "List routers which have all given tag(s) (Comma-separated list of tags)")
	routerCmd.AddCommand(router_listCmd)
}
