package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consumer_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display consumer details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consumer_showCmd).Standalone()

	consumer_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	consumer_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	consumer_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	consumer_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	consumer_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	consumer_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	consumer_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	consumer_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	consumerCmd.AddCommand(consumer_showCmd)
}
