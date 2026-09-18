package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_createCmd).Standalone()

	container_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	container_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	container_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	container_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	container_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	container_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	container_createCmd.Flags().Bool("public", false, "Make the container publicly accessible")
	container_createCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	container_createCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	container_createCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	container_createCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	container_createCmd.Flags().String("storage-policy", "", "Specify a particular storage policy to use.")
	containerCmd.AddCommand(container_createCmd)
}
