package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var snapshot_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a snapshot taken previously with snapshot save",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshot_restoreCmd).Standalone()

	snapshot_restoreCmd.Flags().Bool("no-provision", false, "Disable provisioning")
	snapshot_restoreCmd.Flags().Bool("no-start", false, "Don't start the snapshot after the restore")
	snapshot_restoreCmd.Flags().Bool("provision", false, "Enable provisioning")
	snapshot_restoreCmd.Flags().String("provision-with", "", "Enable only certain provisioners, by type or by name.")
	snapshotCmd.AddCommand(snapshot_restoreCmd)

	carapace.Gen(snapshot_restoreCmd).FlagCompletion(carapace.ActionMap{
		"provision-with": vagrant.ActionProvisioners().UniqueList(","),
	})

	carapace.Gen(snapshot_restoreCmd).PositionalCompletion(
		vagrant.ActionLocalMachines(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionSnapshots(c.Args[0])
		}),
	)
}
