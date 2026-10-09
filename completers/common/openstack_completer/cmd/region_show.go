package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var region_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display region details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(region_showCmd).Standalone()

	region_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	region_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	region_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	region_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	region_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	region_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	region_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	region_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	regionCmd.AddCommand(region_showCmd)
}
