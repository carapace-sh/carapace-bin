package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var consistency_group_add_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Add volume(s) to consistency group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_add_volumeCmd).Standalone()

	consistency_group_addCmd.AddCommand(consistency_group_add_volumeCmd)

	carapace.Gen(consistency_group_add_volumeCmd).PositionalCompletion(
		action.ActionConsistencyGroups(consistency_group_add_volumeCmd),
		action.ActionVolumes(consistency_group_add_volumeCmd),
	)
}
