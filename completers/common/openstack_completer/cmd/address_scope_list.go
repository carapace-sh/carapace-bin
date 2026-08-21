package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_scope_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List address scopes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_scope_listCmd).Standalone()

	address_scope_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	address_scope_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	address_scope_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	address_scope_listCmd.Flags().String("ip-version", "", "List only address scopes with the specified IP version networks (4 or 6)")
	address_scope_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	address_scope_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	address_scope_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	address_scope_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	address_scope_listCmd.Flags().String("name", "", "List only address scopes with the specified name")
	address_scope_listCmd.Flags().Bool("no-share", false, "List only address scopes not shared between projects")
	address_scope_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	address_scope_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	address_scope_listCmd.Flags().String("project", "", "List only address scopes with the specified project (name or ID)")
	address_scope_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	address_scope_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	address_scope_listCmd.Flags().Bool("share", false, "List only address scopes shared between projects")
	address_scope_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	address_scope_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	address_scope_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	address_scopeCmd.AddCommand(address_scope_listCmd)
}
