package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display aggregate details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_showCmd).Standalone()

	aggregate_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	aggregate_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	aggregate_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	aggregate_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	aggregate_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	aggregate_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	aggregate_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	aggregate_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	aggregateCmd.AddCommand(aggregate_showCmd)
}
