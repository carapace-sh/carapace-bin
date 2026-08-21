package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_quota_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Quota",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_quota_deleteCmd).Standalone()

	share_quota_deleteCmd.Flags().String("share-type", "", "Name or ID of a share type to delete the quotas for.")
	share_quota_deleteCmd.Flags().String("user", "", "Name or ID of user to delete the quotas for.")
	share_quotaCmd.AddCommand(share_quota_deleteCmd)
}
