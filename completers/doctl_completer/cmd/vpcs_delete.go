package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var vpcs_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently delete a VPC",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcs_deleteCmd).Standalone()
	vpcs_deleteCmd.Flags().BoolP("force", "f", false, "Delete the VPC without a confirmation prompt")
	vpcsCmd.AddCommand(vpcs_deleteCmd)
}
