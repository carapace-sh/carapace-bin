package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display router details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_showCmd).Standalone()

	router_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	router_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	router_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	router_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	router_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	router_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	router_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	router_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	routerCmd.AddCommand(router_showCmd)
}
