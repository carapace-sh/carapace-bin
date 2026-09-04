package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_add_gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Add router gateway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_add_gatewayCmd).Standalone()

	router_add_gatewayCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	router_add_gatewayCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	router_add_gatewayCmd.Flags().String("fixed-ip", "", "Desired IP and/or subnet (name or ID) on external gateway: subnet=<subnet>,ip-address=<ip-address> (repeat option to set multiple fixed IP addresses)")
	router_add_gatewayCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	router_add_gatewayCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	router_add_gatewayCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	router_add_gatewayCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	router_add_gatewayCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	router_add_gatewayCmd.Flags().String("variable", "", "==SUPPRESS==")
	router_addCmd.AddCommand(router_add_gatewayCmd)
}
