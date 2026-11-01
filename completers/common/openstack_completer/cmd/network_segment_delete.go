package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var network_segment_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete network segment(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_segment_deleteCmd).Standalone()

	network_segmentCmd.AddCommand(network_segment_deleteCmd)

	carapace.Gen(network_segment_deleteCmd).PositionalAnyCompletion(action.ActionNetworkSegments(network_segment_deleteCmd))
}
