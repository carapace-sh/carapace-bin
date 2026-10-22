package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bgp_speaker_remove_networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Remove a network from a BGP speaker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_speaker_remove_networkCmd).Standalone()

	bgp_speaker_removeCmd.AddCommand(bgp_speaker_remove_networkCmd)

	carapace.Gen(bgp_speaker_remove_networkCmd).PositionalCompletion(
		action.ActionBgpSpeakers(bgp_speaker_remove_networkCmd),
		action.ActionNetworks(bgp_speaker_remove_networkCmd),
	)
}
