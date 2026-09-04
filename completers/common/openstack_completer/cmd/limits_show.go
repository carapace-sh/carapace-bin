package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var limits_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show compute and block storage limits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(limits_showCmd).Standalone()

	limits_showCmd.Flags().Bool("absolute", false, "Show absolute limits")
	limits_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	limits_showCmd.Flags().String("domain", "", "==SUPPRESS==")
	limits_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	limits_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	limits_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	limits_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	limits_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	limits_showCmd.Flags().String("project", "", "Show limits for a specific project (name or ID) (only valid with --absolute)")
	limits_showCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	limits_showCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	limits_showCmd.Flags().Bool("rate", false, "Show rate limits.")
	limits_showCmd.Flags().Bool("reserved", false, "Include reservations count (only valid with --absolute)")
	limits_showCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	limits_showCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	limits_showCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	limitsCmd.AddCommand(limits_showCmd)
}
