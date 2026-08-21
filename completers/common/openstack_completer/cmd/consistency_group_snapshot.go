package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consistency_group_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_snapshotCmd).Standalone()

	consistency_groupCmd.AddCommand(consistency_group_snapshotCmd)
}
