package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quota_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List quotas for all projects with non-default quota values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quota_listCmd).Standalone()

	quota_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	quota_listCmd.Flags().Bool("compute", false, "List compute quota")
	quota_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	quota_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	quota_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	quota_listCmd.Flags().Bool("network", false, "List network quota")
	quota_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	quota_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	quota_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	quota_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	quota_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	quota_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	quota_listCmd.Flags().Bool("volume", false, "List volume quota")
	quotaCmd.AddCommand(quota_listCmd)
}
