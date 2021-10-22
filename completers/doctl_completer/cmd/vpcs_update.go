package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var vpcs_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a VPC's configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcs_updateCmd).Standalone()
	vpcs_updateCmd.Flags().Bool("default", false, "The VPC's default state")
	vpcs_updateCmd.Flags().String("description", "", "The VPC's description")
	vpcs_updateCmd.Flags().String("name", "", "The VPC's name")
	vpcsCmd.AddCommand(vpcs_updateCmd)
}
