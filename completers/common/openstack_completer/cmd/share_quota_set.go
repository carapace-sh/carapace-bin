package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_quota_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set Quota for a project, or project/user or project/share-type or a class.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_quota_setCmd).Standalone()

	share_quota_setCmd.Flags().Bool("class", false, "Update class quota to all projects.")
	share_quota_setCmd.Flags().String("encryption-keys", "", "New value for the \"encryption-keys\" quota.")
	share_quota_setCmd.Flags().Bool("force", false, "Force update the quota.")
	share_quota_setCmd.Flags().String("gigabytes", "", "New value for the \"gigabytes\" quota.")
	share_quota_setCmd.Flags().String("per-share-gigabytes", "", "New value for the \"per-share-gigabytes\" quota.")
	share_quota_setCmd.Flags().String("replica-gigabytes", "", "Capacity of share replicas in total.")
	share_quota_setCmd.Flags().String("share-group-snapshots", "", "New value for the \"share-group-snapshots\" quota.")
	share_quota_setCmd.Flags().String("share-groups", "", "New value for the \"share-groups\" quota.")
	share_quota_setCmd.Flags().String("share-networks", "", "New value for the \"share-networks\" quota.")
	share_quota_setCmd.Flags().String("share-replicas", "", "Number of share replicas.")
	share_quota_setCmd.Flags().String("share-type", "", "Name or ID of a share type to set the quotas for.")
	share_quota_setCmd.Flags().String("shares", "", "New value for the \"shares\" quota.")
	share_quota_setCmd.Flags().String("snapshot-gigabytes", "", "New value for the \"snapshot-gigabytes\" quota.")
	share_quota_setCmd.Flags().String("snapshots", "", "New value for the \"snapshots\" quota.")
	share_quota_setCmd.Flags().String("user", "", "Name or ID of a user to set the quotas for.")
	share_quotaCmd.AddCommand(share_quota_setCmd)
}
