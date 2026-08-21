package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extension_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show API extension",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_showCmd).Standalone()

	extension_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	extension_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	extension_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	extension_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	extension_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	extension_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	extension_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	extension_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	extensionCmd.AddCommand(extension_showCmd)
}
