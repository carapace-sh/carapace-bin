package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_service_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tap services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_service_listCmd).Standalone()

	tap_service_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	tap_service_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	tap_service_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	tap_service_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	tap_service_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	tap_service_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	tap_service_listCmd.Flags().String("project", "", "Owner's project (name or ID)")
	tap_service_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	tap_service_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	tap_service_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	tap_service_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	tap_service_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	tap_serviceCmd.AddCommand(tap_service_listCmd)
}
