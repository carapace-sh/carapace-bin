package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List security groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_listCmd).Standalone()

	security_group_listCmd.Flags().String("any-tags", "", "List security group which have any given tag(s) (Comma-separated list of tags)")
	security_group_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	security_group_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	security_group_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	security_group_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	security_group_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	security_group_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	security_group_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	security_group_listCmd.Flags().Bool("no-share", false, "List only security groups not shared between projects")
	security_group_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	security_group_listCmd.Flags().String("not-any-tags", "", "Exclude security group which have any given tag(s) (Comma-separated list of tags)")
	security_group_listCmd.Flags().String("not-tags", "", "Exclude security group which have all given tag(s) (Comma-separated list of tags)")
	security_group_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	security_group_listCmd.Flags().String("project", "", "List only security groups with the specified project (name or ID)")
	security_group_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	security_group_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	security_group_listCmd.Flags().Bool("share", false, "List only security groups shared between projects")
	security_group_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	security_group_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	security_group_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	security_group_listCmd.Flags().String("tags", "", "List security group which have all given tag(s) (Comma-separated list of tags)")
	security_groupCmd.AddCommand(security_group_listCmd)
}
