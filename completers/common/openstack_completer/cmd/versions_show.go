package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versions_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show available versions of services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versions_showCmd).Standalone()

	versions_showCmd.Flags().Bool("all-interfaces", false, "Show values for all interfaces")
	versions_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	versions_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	versions_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	versions_showCmd.Flags().String("interface", "", "Show versions for a specific interface.")
	versions_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	versions_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	versions_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	versions_showCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	versions_showCmd.Flags().String("region-name", "", "Show versions for a specific region.")
	versions_showCmd.Flags().String("service", "", "Show versions for a specific service.")
	versions_showCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	versions_showCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	versions_showCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	versions_showCmd.Flags().String("status", "", "Show versions for a specific status.")
	versionsCmd.AddCommand(versions_showCmd)
}
