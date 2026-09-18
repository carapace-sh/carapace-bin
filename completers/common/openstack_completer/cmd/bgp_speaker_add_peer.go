package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_speaker_add_peerCmd = &cobra.Command{
	Use:   "peer",
	Short: "Add a peer to a BGP speaker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_speaker_add_peerCmd).Standalone()

	bgp_speaker_addCmd.AddCommand(bgp_speaker_add_peerCmd)
}
