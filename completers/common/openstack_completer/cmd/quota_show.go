package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quota_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show quotas for project or class.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quota_showCmd).Standalone()

	quota_showCmd.Flags().Bool("all", false, "Show quotas for all services")
	quota_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	quota_showCmd.Flags().Bool("compute", false, "Show compute quota")
	quota_showCmd.Flags().Bool("default", false, "Show default quotas for <project>")
	quota_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	quota_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	quota_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	quota_showCmd.Flags().Bool("network", false, "Show network quota")
	quota_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	quota_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	quota_showCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	quota_showCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	quota_showCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	quota_showCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	quota_showCmd.Flags().Bool("usage", false, "Show details about quotas usage")
	quota_showCmd.Flags().Bool("volume", false, "Show volume quota")
	quotaCmd.AddCommand(quota_showCmd)
}
