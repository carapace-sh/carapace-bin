package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var user_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List users",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_listCmd).Standalone()

	user_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	user_listCmd.Flags().Bool("disabled", false, "List only disabled users, does nothing with --project and --group")
	user_listCmd.Flags().String("domain", "", "Filter users by <domain> (name or ID)")
	user_listCmd.Flags().Bool("enabled", false, "List only enabled users, does nothing with --project and --group")
	user_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	user_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	user_listCmd.Flags().String("group", "", "Filter users by <group> membership (name or ID)")
	user_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	user_listCmd.Flags().Bool("long", false, "List additional fields in output")
	user_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	user_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	user_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	user_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	user_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	user_listCmd.Flags().String("project", "", "Filter users by <project> (name or ID)")
	user_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	user_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	user_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	user_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	userCmd.AddCommand(user_listCmd)
}
