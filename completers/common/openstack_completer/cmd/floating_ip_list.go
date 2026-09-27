package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List floating IP(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_listCmd).Standalone()

	floating_ip_listCmd.Flags().String("any-tags", "", "List floating IP which have any given tag(s) (Comma-separated list of tags)")
	floating_ip_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	floating_ip_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	floating_ip_listCmd.Flags().String("fixed-ip-address", "", "List only floating IP(s) with the specified fixed IP address")
	floating_ip_listCmd.Flags().String("floating-ip-address", "", "List only floating IP(s) with the specified floating IP address")
	floating_ip_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	floating_ip_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	floating_ip_listCmd.Flags().Bool("long", false, "List additional fields in output")
	floating_ip_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	floating_ip_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	floating_ip_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	floating_ip_listCmd.Flags().String("network", "", "List only floating IP(s) with the specified network (name or ID) (repeat option to fiter on multiple networks)")
	floating_ip_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	floating_ip_listCmd.Flags().String("not-any-tags", "", "Exclude floating IP which have any given tag(s) (Comma-separated list of tags)")
	floating_ip_listCmd.Flags().String("not-tags", "", "Exclude floating IP which have all given tag(s) (Comma-separated list of tags)")
	floating_ip_listCmd.Flags().String("port", "", "List only floating IP(s) with the specified port (name or ID) (repeat option to fiter on multiple ports)")
	floating_ip_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	floating_ip_listCmd.Flags().String("project", "", "List only floating IP(s) with the specified project (name or ID)")
	floating_ip_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	floating_ip_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	floating_ip_listCmd.Flags().String("router", "", "List only floating IP(s) with the specified router (name or ID) (repeat option to fiter on multiple routers)")
	floating_ip_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	floating_ip_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	floating_ip_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	floating_ip_listCmd.Flags().String("status", "", "List only floating IP(s) with the specified status ('ACTIVE', 'DOWN')")
	floating_ip_listCmd.Flags().String("tags", "", "List floating IP which have all given tag(s) (Comma-separated list of tags)")
	floating_ipCmd.AddCommand(floating_ip_listCmd)
}
