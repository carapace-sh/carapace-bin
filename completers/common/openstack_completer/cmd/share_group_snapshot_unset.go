package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_snapshot_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a share group snapshot property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_snapshot_unsetCmd).Standalone()

	share_group_snapshot_unsetCmd.Flags().Bool("description", false, "Unset share group snapshot description.")
	share_group_snapshot_unsetCmd.Flags().Bool("name", false, "Unset share group snapshot name.")
	share_group_snapshotCmd.AddCommand(share_group_snapshot_unsetCmd)
}
