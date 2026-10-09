package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var port_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List ports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(port_listCmd).Standalone()

	port_listCmd.Flags().String("any-tags", "", "List ports which have any given tag(s) (Comma-separated list of tags)")
	port_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	port_listCmd.Flags().String("device-id", "", "List only ports with the specified device ID")
	port_listCmd.Flags().String("device-owner", "", "List only ports with the specified device owner.")
	port_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	port_listCmd.Flags().String("fixed-ip", "", "Desired IP and/or subnet for filtering ports (name or ID): subnet=<subnet>,ip-address=<ip-address>,ip-substring=<ip-substring> (repeat option to filter multiple fixed IP addresses)")
	port_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	port_listCmd.Flags().String("host", "", "List only ports bound to this host ID")
	port_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	port_listCmd.Flags().Bool("long", false, "List additional fields in output")
	port_listCmd.Flags().String("mac-address", "", "List only ports with the specified MAC address")
	port_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	port_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	port_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	port_listCmd.Flags().String("name", "", "List only ports with the specified name")
	port_listCmd.Flags().String("network", "", "List only ports connected to this network (name or ID)")
	port_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	port_listCmd.Flags().String("not-any-tags", "", "Exclude ports which have any given tag(s) (Comma-separated list of tags)")
	port_listCmd.Flags().String("not-tags", "", "Exclude ports which have all given tag(s) (Comma-separated list of tags)")
	port_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	port_listCmd.Flags().String("project", "", "List only ports with the specified project (name or ID)")
	port_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	port_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	port_listCmd.Flags().String("router", "", "List only ports attached to this router (name or ID)")
	port_listCmd.Flags().String("security-group", "", "List only ports associated with this security group")
	port_listCmd.Flags().String("server", "", "List only ports attached to this server (name or ID)")
	port_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	port_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	port_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	port_listCmd.Flags().String("status", "", "List only ports with the specified status ('ACTIVE', 'BUILD', 'DOWN', 'ERROR')")
	port_listCmd.Flags().String("tags", "", "List ports which have all given tag(s) (Comma-separated list of tags)")
	portCmd.AddCommand(port_listCmd)
}
