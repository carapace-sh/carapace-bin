package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:     "rollback [-rRf] snapshot",
	Short:   "rollback to a snapshot",
	GroupID: "dataset",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollbackCmd).Standalone()

	rollbackCmd.Flags().BoolS("R", "R", false, "destroy later snapshots and clones")
	rollbackCmd.Flags().BoolS("f", "f", false, "force unmount")
	rollbackCmd.Flags().BoolS("r", "r", false, "destroy later snapshots")

	rootCmd.AddCommand(rollbackCmd)

	carapace.Gen(rollbackCmd).PositionalCompletion(
		zfs.ActionSnapshots(),
	)
}
