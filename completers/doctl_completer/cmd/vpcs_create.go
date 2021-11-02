package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var vpcs_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new VPC",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcs_createCmd).Standalone()
	vpcs_createCmd.Flags().String("description", "", "The VPC's name")
	vpcs_createCmd.Flags().String("ip-range", "", "The range of IP addresses in the VPC in CIDR notation, e.g.: `10.116.0.0/20`")
	vpcs_createCmd.Flags().String("name", "", "The VPC's name (required)")
	vpcs_createCmd.Flags().String("region", "", "The VPC's region slug, e.g.: `nyc1` (required)")
	vpcsCmd.AddCommand(vpcs_createCmd)
}
