package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_remove_hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Remove host from aggregate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_remove_hostCmd).Standalone()

	aggregate_remove_hostCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	aggregate_remove_hostCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	aggregate_remove_hostCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	aggregate_remove_hostCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	aggregate_remove_hostCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	aggregate_remove_hostCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	aggregate_remove_hostCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	aggregate_remove_hostCmd.Flags().String("variable", "", "==SUPPRESS==")
	aggregate_removeCmd.AddCommand(aggregate_remove_hostCmd)
}
