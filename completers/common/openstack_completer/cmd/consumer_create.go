package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consumer_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new consumer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consumer_createCmd).Standalone()

	consumer_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	consumer_createCmd.Flags().String("description", "", "New consumer description")
	consumer_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	consumer_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	consumer_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	consumer_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	consumer_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	consumer_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	consumer_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	consumerCmd.AddCommand(consumer_createCmd)
}
