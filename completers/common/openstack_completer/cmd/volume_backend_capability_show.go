package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backend_capability_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show capability command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backend_capability_showCmd).Standalone()

	volume_backend_capability_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_backend_capability_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_backend_capability_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_backend_capability_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_backend_capability_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_backend_capability_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_backend_capability_showCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	volume_backend_capability_showCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	volume_backend_capability_showCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	volume_backend_capability_showCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	volume_backend_capabilityCmd.AddCommand(volume_backend_capability_showCmd)
}
