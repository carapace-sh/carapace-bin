package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var snapshot_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snapshots taken for a machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_listCmd).Standalone()

	snapshotCmd.AddCommand(snapshot_listCmd)

	carapace.Gen(snapshot_listCmd).PositionalCompletion(
		vagrant.ActionLocalMachines(),
	)
}
