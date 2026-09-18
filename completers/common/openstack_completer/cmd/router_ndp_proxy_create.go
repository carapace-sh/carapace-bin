package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_ndp_proxy_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create NDP proxy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_ndp_proxy_createCmd).Standalone()

	router_ndp_proxy_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	router_ndp_proxy_createCmd.Flags().String("description", "", "Text to describe/contextualize the use of the NDP proxy configuration")
	router_ndp_proxy_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	router_ndp_proxy_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	router_ndp_proxy_createCmd.Flags().String("ip-address", "", "The IPv6 address that is to be proxied.")
	router_ndp_proxy_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	router_ndp_proxy_createCmd.Flags().String("name", "", "New NDP proxy name")
	router_ndp_proxy_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	router_ndp_proxy_createCmd.Flags().String("port", "", "The name or ID of the network port associated to the NDP proxy")
	router_ndp_proxy_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	router_ndp_proxy_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	router_ndp_proxy_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	router_ndp_proxy_createCmd.MarkFlagRequired("port")
	router_ndp_proxyCmd.AddCommand(router_ndp_proxy_createCmd)
}
