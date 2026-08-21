package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_remove_routeCmd = &cobra.Command{
	Use:   "route",
	Short: "Remove extra static routes from a router's routing table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_remove_routeCmd).Standalone()

	router_remove_routeCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	router_remove_routeCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	router_remove_routeCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	router_remove_routeCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	router_remove_routeCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	router_remove_routeCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	router_remove_routeCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	router_remove_routeCmd.Flags().String("route", "", "Remove extra static route from the router.")
	router_remove_routeCmd.Flags().String("variable", "", "==SUPPRESS==")
	router_removeCmd.AddCommand(router_remove_routeCmd)
}
