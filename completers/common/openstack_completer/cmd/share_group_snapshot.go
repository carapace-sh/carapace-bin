package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_snapshotCmd).Standalone()

	share_groupCmd.AddCommand(share_group_snapshotCmd)
}
