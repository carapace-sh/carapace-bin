package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_network_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List share networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_network_listCmd).Standalone()

	share_network_listCmd.Flags().Bool("all-projects", false, "Include all projects (admin only)")
	share_network_listCmd.Flags().String("cidr", "", "Filter share networks by the CIDR of network.")
	share_network_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_network_listCmd.Flags().String("created-before", "", "Filter share networks by date they were created before.")
	share_network_listCmd.Flags().String("created-since", "", "Filter share networks by date they were created after.")
	share_network_listCmd.Flags().String("description", "", "Filter share networks by description.")
	share_network_listCmd.Flags().String("description~", "", "Filter share networks by description-pattern.")
	share_network_listCmd.Flags().Bool("detail", false, "List share networks with details")
	share_network_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_network_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_network_listCmd.Flags().String("ip-version", "", "Filter share networks by the IP Version of the network, either 4 or 6.")
	share_network_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_network_listCmd.Flags().String("name", "", "Filter share networks by name")
	share_network_listCmd.Flags().String("name~", "", "Filter share networks by name-pattern.")
	share_network_listCmd.Flags().String("network-type", "", "Filter share networks by the type of network.")
	share_network_listCmd.Flags().String("neutron-net-id", "", "Filter share networks by the ID of a neutron network.")
	share_network_listCmd.Flags().String("neutron-subnet-id", "", "Filter share networks by the ID of a neutron sub network.")
	share_network_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_network_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_network_listCmd.Flags().String("project", "", "Filter share networks by project (name or ID) (admin only)")
	share_network_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	share_network_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_network_listCmd.Flags().String("security-service", "", "Filter share networks by the name or ID of a security service attached to the network.")
	share_network_listCmd.Flags().String("segmentation-id", "", "Filter share networks by the segmentation ID of network.")
	share_network_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_network_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_network_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_networkCmd.AddCommand(share_network_listCmd)
}
