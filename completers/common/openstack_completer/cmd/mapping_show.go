package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mapping_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display mapping details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mapping_showCmd).Standalone()

	mapping_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	mapping_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	mapping_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	mapping_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	mapping_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	mapping_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	mapping_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	mapping_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	mappingCmd.AddCommand(mapping_showCmd)
}
