package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_endpoint_group_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an endpoint group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_endpoint_group_createCmd).Standalone()

	vpn_endpoint_group_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_endpoint_group_createCmd.Flags().String("description", "", "Description for the endpoint group")
	vpn_endpoint_group_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_endpoint_group_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_endpoint_group_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_endpoint_group_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_endpoint_group_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	vpn_endpoint_group_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_endpoint_group_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	vpn_endpoint_group_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	vpn_endpoint_group_createCmd.Flags().String("type", "", "Type of endpoints in group (e.g. subnet, cidr, network, router).")
	vpn_endpoint_group_createCmd.Flags().String("value", "", "Endpoint(s) for the group.")
	vpn_endpoint_group_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	vpn_endpoint_group_createCmd.MarkFlagRequired("type")
	vpn_endpoint_group_createCmd.MarkFlagRequired("value")
	vpn_endpoint_groupCmd.AddCommand(vpn_endpoint_group_createCmd)
}
