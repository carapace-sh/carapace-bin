package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subnet_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List subnets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_listCmd).Standalone()

	subnet_listCmd.Flags().String("any-tags", "", "List subnets which have any given tag(s) (Comma-separated list of tags)")
	subnet_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	subnet_listCmd.Flags().Bool("dhcp", false, "List only subnets which have DHCP enabled")
	subnet_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	subnet_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	subnet_listCmd.Flags().String("gateway", "", "List only subnets with the specified gateway IP")
	subnet_listCmd.Flags().String("ip-version", "", "List only subnets with the specified IP version.")
	subnet_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	subnet_listCmd.Flags().Bool("long", false, "List additional fields in output")
	subnet_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	subnet_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	subnet_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	subnet_listCmd.Flags().String("name", "", "List only subnets with the specified name")
	subnet_listCmd.Flags().String("network", "", "List only subnets which belong to the specified network (name or ID)")
	subnet_listCmd.Flags().Bool("no-dhcp", false, "List only subnets which have DHCP disabled")
	subnet_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	subnet_listCmd.Flags().String("not-any-tags", "", "Exclude subnets which have any given tag(s) (Comma-separated list of tags)")
	subnet_listCmd.Flags().String("not-tags", "", "Exclude subnets which have all given tag(s) (Comma-separated list of tags)")
	subnet_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	subnet_listCmd.Flags().String("project", "", "List only subnets with the specified project (name or ID)")
	subnet_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	subnet_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	subnet_listCmd.Flags().String("service-type", "", "List only subnets with the specified service type, for example, network:floatingip_agent_gateway.")
	subnet_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	subnet_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	subnet_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	subnet_listCmd.Flags().String("subnet-pool", "", "List only subnets which belong to the specified subnet pool (name or ID)")
	subnet_listCmd.Flags().String("subnet-range", "", "List only subnets with the specified subnet range (in CIDR notation).")
	subnet_listCmd.Flags().String("tags", "", "List subnets which have all given tag(s) (Comma-separated list of tags)")
	subnetCmd.AddCommand(subnet_listCmd)
}
