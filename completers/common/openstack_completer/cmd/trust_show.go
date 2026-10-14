package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trust_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display trust details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_showCmd).Standalone()

	trust_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	trust_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	trust_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	trust_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	trust_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	trust_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	trust_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	trust_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	trustCmd.AddCommand(trust_showCmd)
}
