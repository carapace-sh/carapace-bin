package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_type_access_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deny a project to access a share group type (Admin only).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_type_access_deleteCmd).Standalone()

	share_group_type_access_deleteCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	share_group_type_accessCmd.AddCommand(share_group_type_access_deleteCmd)
}
