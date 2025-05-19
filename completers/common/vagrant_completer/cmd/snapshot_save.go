package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var snapshot_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Take a snapshot of the current state of the machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_saveCmd).Standalone()

	snapshot_saveCmd.Flags().BoolP("force", "f", false, "Replace snapshot without confirmation")
	snapshotCmd.AddCommand(snapshot_saveCmd)

	carapace.Gen(snapshot_saveCmd).PositionalCompletion(
		vagrant.ActionLocalMachines(),
	)
}
