package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_network_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a share network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_network_createCmd).Standalone()

	share_network_createCmd.Flags().String("availability-zone", "", "Name or ID of the avalilability zone to assign the specified network subnet parameters to.")
	share_network_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_network_createCmd.Flags().String("description", "", "Add a description to the share network (Optional).")
	share_network_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_network_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_network_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_network_createCmd.Flags().String("name", "", "Add a name to the share network (Optional)")
	share_network_createCmd.Flags().String("neutron-net-id", "", "ID of the neutron network that must be associated with the share network (Optional).")
	share_network_createCmd.Flags().String("neutron-subnet-id", "", "ID of the neutron sub-network that must be associated with the share network (Optional).")
	share_network_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_network_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_network_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_network_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_networkCmd.AddCommand(share_network_createCmd)
}
