package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new aggregate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_createCmd).Standalone()

	aggregate_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	aggregate_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	aggregate_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	aggregate_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	aggregate_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	aggregate_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	aggregate_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	aggregate_createCmd.Flags().String("property", "", "Property to add to this aggregate (repeat option to set multiple properties)")
	aggregate_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	aggregate_createCmd.Flags().String("zone", "", "Availability zone name")
	aggregateCmd.AddCommand(aggregate_createCmd)
}
