package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_snapshotCmd).Standalone()

	volume_groupCmd.AddCommand(volume_group_snapshotCmd)
}
