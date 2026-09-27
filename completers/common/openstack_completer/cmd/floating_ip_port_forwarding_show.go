package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_port_forwarding_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display floating IP Port Forwarding details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_port_forwarding_showCmd).Standalone()

	floating_ip_port_forwarding_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	floating_ip_port_forwarding_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	floating_ip_port_forwarding_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	floating_ip_port_forwarding_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	floating_ip_port_forwarding_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	floating_ip_port_forwarding_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	floating_ip_port_forwarding_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	floating_ip_port_forwarding_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	floating_ip_port_forwardingCmd.AddCommand(floating_ip_port_forwarding_showCmd)
}
