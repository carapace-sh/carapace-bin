package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_snapshot_membersCmd = &cobra.Command{
	Use:   "members",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_snapshot_membersCmd).Standalone()

	share_group_snapshotCmd.AddCommand(share_group_snapshot_membersCmd)
}
