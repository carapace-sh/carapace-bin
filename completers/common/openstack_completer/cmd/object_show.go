package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var object_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display object details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_showCmd).Standalone()

	object_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	object_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	object_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	object_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	object_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	object_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	object_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	object_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	objectCmd.AddCommand(object_showCmd)
}
