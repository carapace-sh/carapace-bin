package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_ndp_proxy_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List NDP proxies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_ndp_proxy_listCmd).Standalone()

	router_ndp_proxy_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	router_ndp_proxy_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	router_ndp_proxy_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	router_ndp_proxy_listCmd.Flags().String("ip-address", "", "List only NDP proxies associated with the specified IPv6 address")
	router_ndp_proxy_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	router_ndp_proxy_listCmd.Flags().String("name", "", "List only NDP proxies with the specified name")
	router_ndp_proxy_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	router_ndp_proxy_listCmd.Flags().String("port", "", "List only NDP proxies associated with the specified port (name or ID)")
	router_ndp_proxy_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	router_ndp_proxy_listCmd.Flags().String("project", "", "List only NDP proxies with the specified project (name or ID)")
	router_ndp_proxy_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	router_ndp_proxy_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	router_ndp_proxy_listCmd.Flags().String("router", "", "List only NDP proxies associated with the specifed router (name or ID)")
	router_ndp_proxy_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	router_ndp_proxy_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	router_ndp_proxy_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	router_ndp_proxyCmd.AddCommand(router_ndp_proxy_listCmd)
}
