package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_network_subnet_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a share network subnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_network_subnet_createCmd).Standalone()

	share_network_subnet_createCmd.Flags().String("availability-zone", "", "Optional availability zone that the subnet is available within (Default=None).")
	share_network_subnet_createCmd.Flags().Bool("check-only", false, "Run a dry-run of a share network subnet create.")
	share_network_subnet_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_network_subnet_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_network_subnet_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_network_subnet_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_network_subnet_createCmd.Flags().String("neutron-net-id", "", "Neutron network ID.")
	share_network_subnet_createCmd.Flags().String("neutron-subnet-id", "", "Neutron subnet ID.")
	share_network_subnet_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_network_subnet_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_network_subnet_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_network_subnet_createCmd.Flags().String("property", "", "Set a property to this share network subnet (repeat option to set multiple properties).")
	share_network_subnet_createCmd.Flags().Bool("restart-check", false, "Restart a dry-run of a share network subnet create.")
	share_network_subnet_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_network_subnetCmd.AddCommand(share_network_subnet_createCmd)
}
