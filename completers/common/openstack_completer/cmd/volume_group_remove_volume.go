package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var volume_group_remove_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Remove volume(s) from volume group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_remove_volumeCmd).Standalone()

	volume_group_removeCmd.AddCommand(volume_group_remove_volumeCmd)

	carapace.Gen(volume_group_remove_volumeCmd).PositionalCompletion(
		action.ActionVolumeGroups(volume_group_remove_volumeCmd),
		action.ActionVolumes(volume_group_remove_volumeCmd),
	)
}
