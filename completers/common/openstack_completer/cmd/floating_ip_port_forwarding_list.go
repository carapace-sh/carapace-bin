package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_port_forwarding_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List floating IP port forwarding",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_port_forwarding_listCmd).Standalone()

	floating_ip_port_forwarding_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	floating_ip_port_forwarding_listCmd.Flags().String("external-protocol-port", "", "List only floating IP port forwardings with the specified external protocol port number")
	floating_ip_port_forwarding_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	floating_ip_port_forwarding_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	floating_ip_port_forwarding_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	floating_ip_port_forwarding_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	floating_ip_port_forwarding_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	floating_ip_port_forwarding_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	floating_ip_port_forwarding_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	floating_ip_port_forwarding_listCmd.Flags().String("port", "", "List only floating IP port forwardings with the specified internal network port (name or ID)")
	floating_ip_port_forwarding_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	floating_ip_port_forwarding_listCmd.Flags().String("protocol", "", "List only floating IP port forwardings with the specified protocol number")
	floating_ip_port_forwarding_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	floating_ip_port_forwarding_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	floating_ip_port_forwarding_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	floating_ip_port_forwarding_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	floating_ip_port_forwardingCmd.AddCommand(floating_ip_port_forwarding_listCmd)
}
