package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bgp_speaker_remove_peerCmd = &cobra.Command{
	Use:   "peer",
	Short: "Remove a peer from a BGP speaker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_speaker_remove_peerCmd).Standalone()

	bgp_speaker_removeCmd.AddCommand(bgp_speaker_remove_peerCmd)

	carapace.Gen(bgp_speaker_remove_peerCmd).PositionalCompletion(
		action.ActionBgpSpeakers(),
		action.ActionBgpPeers(),
	)
}
