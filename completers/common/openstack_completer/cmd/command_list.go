package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var command_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List recognized commands by group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(command_listCmd).Standalone()

	command_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	command_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	command_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	command_listCmd.Flags().String("group", "", "Show commands filtered by a command group, for example: identity, volume, compute, image, network and other keywords")
	command_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	command_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	command_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	command_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	command_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	command_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	command_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	commandCmd.AddCommand(command_listCmd)
}
