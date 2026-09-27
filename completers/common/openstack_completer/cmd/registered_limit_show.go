package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registered_limit_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display registered limit details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registered_limit_showCmd).Standalone()

	registered_limit_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	registered_limit_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	registered_limit_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	registered_limit_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	registered_limit_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	registered_limit_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	registered_limit_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	registered_limit_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	registered_limitCmd.AddCommand(registered_limit_showCmd)
}
