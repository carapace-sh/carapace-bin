package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var console_connection_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show server's remote console connection information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(console_connection_showCmd).Standalone()

	console_connection_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	console_connection_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	console_connection_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	console_connection_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	console_connection_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	console_connection_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	console_connection_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	console_connection_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	console_connectionCmd.AddCommand(console_connection_showCmd)
}
