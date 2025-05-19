package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var snapshot_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a snapshot taken previously with snapshot save",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_deleteCmd).Standalone()

	snapshotCmd.AddCommand(snapshot_deleteCmd)

	carapace.Gen(snapshot_deleteCmd).PositionalCompletion(
		vagrant.ActionLocalMachines(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionSnapshots(c.Args[0])
		}),
	)
}
