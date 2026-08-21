package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_listCmd).Standalone()

	group_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	group_listCmd.Flags().String("domain", "", "Filter group list by <domain> (name or ID)")
	group_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	group_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	group_listCmd.Flags().Bool("long", false, "List additional fields in output")
	group_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	group_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	group_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	group_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	group_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	group_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	group_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	group_listCmd.Flags().String("user", "", "Filter group list by <user> (name or ID)")
	group_listCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	groupCmd.AddCommand(group_listCmd)
}
