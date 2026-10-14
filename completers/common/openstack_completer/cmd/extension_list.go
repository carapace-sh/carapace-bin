package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extension_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List API extensions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_listCmd).Standalone()

	extension_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	extension_listCmd.Flags().Bool("compute", false, "List extensions for the Compute API")
	extension_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	extension_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	extension_listCmd.Flags().Bool("identity", false, "List extensions for the Identity API (only supported by v2)")
	extension_listCmd.Flags().Bool("long", false, "List additional fields in output")
	extension_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	extension_listCmd.Flags().Bool("network", false, "List extensions for the Network API")
	extension_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	extension_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	extension_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	extension_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	extension_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	extension_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	extension_listCmd.Flags().Bool("volume", false, "List extensions for the Block Storage API")
	extensionCmd.AddCommand(extension_listCmd)
}
