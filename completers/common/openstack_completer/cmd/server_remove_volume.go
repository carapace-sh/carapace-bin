package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var server_remove_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Remove volume from server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_volumeCmd).Standalone()

	server_removeCmd.AddCommand(server_remove_volumeCmd)

	carapace.Gen(server_remove_volumeCmd).PositionalCompletion(
		action.ActionServers(server_remove_volumeCmd),
		action.ActionVolumes(server_remove_volumeCmd),
	)
}
