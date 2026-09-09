package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_add_hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Add host to aggregate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_add_hostCmd).Standalone()

	aggregate_add_hostCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	aggregate_add_hostCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	aggregate_add_hostCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	aggregate_add_hostCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	aggregate_add_hostCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	aggregate_add_hostCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	aggregate_add_hostCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	aggregate_add_hostCmd.Flags().String("variable", "", "==SUPPRESS==")
	aggregate_addCmd.AddCommand(aggregate_add_hostCmd)
}
