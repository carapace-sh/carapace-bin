package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var routeListCmd = &cobra.Command{
	Use:   "route:list [--json] [--method [METHOD]] [--name [NAME]] [--domain [DOMAIN]] [--path [PATH]] [--except-path [EXCEPT-PATH]] [-r|--reverse] [--sort [SORT]] [--except-vendor] [--only-vendor]",
	Short: "List all registered routes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(routeListCmd).Standalone()

	routeListCmd.Flags().String("domain", "", "Filter the routes by domain")
	routeListCmd.Flags().String("except-path", "", "Do not display the routes matching the given path pattern")
	routeListCmd.Flags().Bool("except-vendor", false, "Do not display routes defined by vendor packages")
	routeListCmd.Flags().Bool("json", false, "Output the route list as JSON")
	routeListCmd.Flags().String("method", "", "Filter the routes by method")
	routeListCmd.Flags().String("name", "", "Filter the routes by name")
	routeListCmd.Flags().Bool("only-vendor", false, "Only display routes defined by vendor packages")
	routeListCmd.Flags().String("path", "", "Only show routes matching the given path pattern")
	routeListCmd.Flags().Bool("reverse", false, "Reverse the ordering of the routes")
	routeListCmd.Flags().String("sort", "", "The column (domain, method, uri, name, action, middleware) to sort by")
	rootCmd.AddCommand(routeListCmd)
}
