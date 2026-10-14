package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_volume_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the volumes attached to a server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_volume_listCmd).Standalone()

	server_volume_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_volume_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_volume_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_volume_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_volume_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_volume_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_volume_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	server_volume_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	server_volume_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	server_volume_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	server_volumeCmd.AddCommand(server_volume_listCmd)
}
