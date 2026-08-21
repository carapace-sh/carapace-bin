package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List servers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_listCmd).Standalone()

	server_listCmd.Flags().Bool("all-projects", false, "Include all projects (admin only) (can be specified using the ALL_PROJECTS envvar)")
	server_listCmd.Flags().String("availability-zone", "", "Search by availability zone (admin only before microversion 2.83)")
	server_listCmd.Flags().String("changes-before", "", "List only servers changed before a certain point of time.")
	server_listCmd.Flags().String("changes-since", "", "List only servers changed after a certain point of time.")
	server_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_listCmd.Flags().Bool("config-drive", false, "Only display servers with a config drive attached (admin only before microversion 2.83)")
	server_listCmd.Flags().Bool("deleted", false, "Only display deleted servers (admin only)")
	server_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_listCmd.Flags().String("flavor", "", "Search by flavor (name or ID)")
	server_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_listCmd.Flags().String("host", "", "Search by hostname")
	server_listCmd.Flags().String("image", "", "Search by image (name or ID)")
	server_listCmd.Flags().String("instance-name", "", "Regular expression to match instance name (admin only)")
	server_listCmd.Flags().String("ip", "", "Regular expression to match IP addresses")
	server_listCmd.Flags().String("ip6", "", "Regular expression to match IPv6 addresses.")
	server_listCmd.Flags().String("key-name", "", "Search by keypair name (admin only before microversion 2.83)")
	server_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	server_listCmd.Flags().Bool("locked", false, "Only display locked servers (supported by --os-compute-api-version 2.73 or above)")
	server_listCmd.Flags().Bool("long", false, "List additional fields in output")
	server_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	server_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	server_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_listCmd.Flags().String("name", "", "Regular expression to match names")
	server_listCmd.Flags().Bool("name-lookup-one-by-one", false, "When looking up flavor and image names, look them up one by one as needed instead of all together (default).")
	server_listCmd.Flags().Bool("no-config-drive", false, "Only display servers without a config drive attached (admin only before microversion 2.83)")
	server_listCmd.Flags().BoolP("no-name-lookup", "n", false, "Skip flavor and image name lookup.")
	server_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_listCmd.Flags().String("not-tags", "", "Only list servers without the specified tag.")
	server_listCmd.Flags().String("power-state", "", "Search by power_state value (admin only before microversion 2.83)")
	server_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_listCmd.Flags().String("progress", "", "Search by progress value (%%) (admin only before microversion 2.83)")
	server_listCmd.Flags().String("project", "", "Search by project (admin only) (name or ID)")
	server_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	server_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	server_listCmd.Flags().String("reservation-id", "", "Only return instances that match the reservation")
	server_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	server_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	server_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	server_listCmd.Flags().String("status", "", "Search by server status")
	server_listCmd.Flags().String("tags", "", "Only list servers with the specified tag.")
	server_listCmd.Flags().String("task-state", "", "Search by task_state value (admin only before microversion 2.83)")
	server_listCmd.Flags().Bool("unlocked", false, "Only display unlocked servers (supported by --os-compute-api-version 2.73 or above)")
	server_listCmd.Flags().String("user", "", "Search by user (name or ID) (admin only before microversion 2.83)")
	server_listCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	server_listCmd.Flags().String("vm-state", "", "Search by vm_state value (admin only before microversion 2.83)")
	serverCmd.AddCommand(server_listCmd)
}
