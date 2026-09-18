package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_instanceCmd).Standalone()

	share_snapshotCmd.AddCommand(share_snapshot_instanceCmd)
}
