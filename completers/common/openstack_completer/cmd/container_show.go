package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display container details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_showCmd).Standalone()

	container_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	container_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	container_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	container_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	container_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	container_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	container_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	container_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	containerCmd.AddCommand(container_showCmd)
}
