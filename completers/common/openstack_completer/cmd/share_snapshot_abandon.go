package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_abandonCmd = &cobra.Command{
	Use:   "abandon",
	Short: "Abandon share snapshot(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_abandonCmd).Standalone()

	share_snapshot_abandonCmd.Flags().Bool("wait", false, "Wait until share snapshot is abandoned")
	share_snapshotCmd.AddCommand(share_snapshot_abandonCmd)
}
