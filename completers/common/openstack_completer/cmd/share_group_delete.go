package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more share groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_deleteCmd).Standalone()

	share_group_deleteCmd.Flags().Bool("force", false, "Attempt to force delete the share group (Default=False) (Admin only).")
	share_group_deleteCmd.Flags().Bool("wait", false, "Wait for share group to delete")
	share_groupCmd.AddCommand(share_group_deleteCmd)
}
