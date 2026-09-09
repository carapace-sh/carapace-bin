package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var local_ip_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Local IPs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(local_ip_listCmd).Standalone()

	local_ip_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	local_ip_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	local_ip_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	local_ip_listCmd.Flags().String("ip-mode", "", "List only local IP(s) with the specified IP mode")
	local_ip_listCmd.Flags().String("local-ip-address", "", "List only local IP(s) with the specified IP address")
	local_ip_listCmd.Flags().String("local-port", "", "List only local IP(s) with the specified port (name or ID)")
	local_ip_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	local_ip_listCmd.Flags().String("name", "", "List only local IP(s) with the specified name")
	local_ip_listCmd.Flags().String("network", "", "List only local IP(s) with the specified network (name or ID)")
	local_ip_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	local_ip_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	local_ip_listCmd.Flags().String("project", "", "List only local IP(s) with the specified project (name or ID)")
	local_ip_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	local_ip_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	local_ip_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	local_ip_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	local_ip_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	local_ipCmd.AddCommand(local_ip_listCmd)
}
