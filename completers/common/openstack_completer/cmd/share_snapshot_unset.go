package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a share snapshot property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_unsetCmd).Standalone()

	share_snapshot_unsetCmd.Flags().Bool("description", false, "Unset snapshot description.")
	share_snapshot_unsetCmd.Flags().Bool("name", false, "Unset snapshot name.")
	share_snapshot_unsetCmd.Flags().String("property", "", "Remove a property from snapshot (repeat option to remove multiple properties)")
	share_snapshotCmd.AddCommand(share_snapshot_unsetCmd)
}
