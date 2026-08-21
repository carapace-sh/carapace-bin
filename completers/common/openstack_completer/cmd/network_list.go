package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_listCmd).Standalone()

	network_listCmd.Flags().String("agent", "", "List only networks hosted the specified agent (ID only)")
	network_listCmd.Flags().String("any-tags", "", "List networks which have any given tag(s) (Comma-separated list of tags)")
	network_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_listCmd.Flags().Bool("disable", false, "List only disabled networks")
	network_listCmd.Flags().Bool("enable", false, "List only enabled networks")
	network_listCmd.Flags().Bool("external", false, "List only external networks")
	network_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_listCmd.Flags().Bool("internal", false, "List only internal networks")
	network_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	network_listCmd.Flags().Bool("long", false, "List additional fields in output")
	network_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	network_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	network_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_listCmd.Flags().String("name", "", "List only networks with the specified name")
	network_listCmd.Flags().Bool("no-share", false, "List only networks not shared between projects")
	network_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_listCmd.Flags().String("not-any-tags", "", "Exclude networks which have any given tag(s) (Comma-separated list of tags)")
	network_listCmd.Flags().String("not-tags", "", "Exclude networks which have all given tag(s) (Comma-separated list of tags)")
	network_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_listCmd.Flags().String("project", "", "List only networks with the specified project (name or ID)")
	network_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_listCmd.Flags().String("provider-network-type", "", "List only networks with the specified physical mechanisms.")
	network_listCmd.Flags().String("provider-physical-network", "", "List only networks with the specified physical network name")
	network_listCmd.Flags().String("provider-segment", "", "List only networks with the specified provider segment ID (VLAN ID for VLAN networks or Tunnel ID for GENEVE/GRE/VXLAN networks)")
	network_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_listCmd.Flags().Bool("share", false, "List only networks shared between projects")
	network_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_listCmd.Flags().String("status", "", "List only networks with the specified status ('ACTIVE', 'BUILD', 'DOWN', 'ERROR')")
	network_listCmd.Flags().String("tags", "", "List networks which have all given tag(s) (Comma-separated list of tags)")
	networkCmd.AddCommand(network_listCmd)
}
