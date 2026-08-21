package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var project_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_listCmd).Standalone()

	project_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	project_listCmd.Flags().Bool("disabled", false, "List only disabled projects")
	project_listCmd.Flags().String("domain", "", "Filter projects by <domain> (name or ID)")
	project_listCmd.Flags().Bool("enabled", false, "List only enabled projects")
	project_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	project_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	project_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	project_listCmd.Flags().Bool("long", false, "List additional fields in output")
	project_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	project_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	project_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	project_listCmd.Flags().Bool("my-projects", false, "List projects for the authenticated user.")
	project_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	project_listCmd.Flags().String("not-tags", "", "Exclude projects which have all given tag(s) (Comma-separated list of tags)")
	project_listCmd.Flags().String("not-tags-any", "", "Exclude projects which have any given tag(s) (Comma-separated list of tags)")
	project_listCmd.Flags().String("parent", "", "Filter projects whose parent is <parent> (name or ID)")
	project_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	project_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	project_listCmd.Flags().String("sort", "", "Sort output by selected keys and directions (asc or desc) (default: asc), repeat this option to specify multiple keys and directions.")
	project_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	project_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	project_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	project_listCmd.Flags().String("tags", "", "List projects which have all given tag(s) (Comma-separated list of tags)")
	project_listCmd.Flags().String("tags-any", "", "List projects which have any given tag(s) (Comma-separated list of tags)")
	project_listCmd.Flags().String("user", "", "Filter projects by <user> (name or ID)")
	projectCmd.AddCommand(project_listCmd)
}
