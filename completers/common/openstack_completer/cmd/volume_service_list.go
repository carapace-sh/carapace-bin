package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_service_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List service command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_service_listCmd).Standalone()

	volume_service_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_service_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_service_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_service_listCmd.Flags().String("host", "", "List services on specified host (name only)")
	volume_service_listCmd.Flags().Bool("long", false, "List additional fields in output")
	volume_service_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_service_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_service_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_service_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	volume_service_listCmd.Flags().String("service", "", "List only specified service (name only)")
	volume_service_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	volume_service_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	volume_service_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	volume_serviceCmd.AddCommand(volume_service_listCmd)
}
