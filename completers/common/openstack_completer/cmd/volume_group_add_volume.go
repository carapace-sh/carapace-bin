package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var volume_group_add_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Add volume(s) to volume group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_add_volumeCmd).Standalone()

	volume_group_addCmd.AddCommand(volume_group_add_volumeCmd)

	carapace.Gen(volume_group_add_volumeCmd).PositionalCompletion(
		action.ActionVolumeGroups(volume_group_add_volumeCmd),
		action.ActionVolumes(volume_group_add_volumeCmd),
	)
}
