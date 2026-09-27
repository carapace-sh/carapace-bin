package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_service_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an VPN service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_service_createCmd).Standalone()

	vpn_service_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_service_createCmd.Flags().String("description", "", "Description for the VPN service")
	vpn_service_createCmd.Flags().Bool("disable", false, "Disable VPN service")
	vpn_service_createCmd.Flags().Bool("enable", false, "Enable VPN service")
	vpn_service_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_service_createCmd.Flags().String("flavor", "", "Flavor for the VPN service (name or ID)")
	vpn_service_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_service_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_service_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_service_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	vpn_service_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_service_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	vpn_service_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	vpn_service_createCmd.Flags().String("router", "", "Router for the VPN service (name or ID)")
	vpn_service_createCmd.Flags().String("subnet", "", "Local private subnet (name or ID)")
	vpn_service_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	vpn_service_createCmd.MarkFlagRequired("router")
	vpn_serviceCmd.AddCommand(vpn_service_createCmd)
}
