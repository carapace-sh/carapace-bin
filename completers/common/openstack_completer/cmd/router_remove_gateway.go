package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_remove_gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Remove router gateway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_remove_gatewayCmd).Standalone()

	router_remove_gatewayCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	router_remove_gatewayCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	router_remove_gatewayCmd.Flags().String("fixed-ip", "", "IP and/or subnet (name or ID) on the external gateway which is used to identify a particular gateway if multiple are attached to the same network: subnet=<subnet>,ip-address=<ip-address>")
	router_remove_gatewayCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	router_remove_gatewayCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	router_remove_gatewayCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	router_remove_gatewayCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	router_remove_gatewayCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	router_remove_gatewayCmd.Flags().String("variable", "", "==SUPPRESS==")
	router_removeCmd.AddCommand(router_remove_gatewayCmd)
}
