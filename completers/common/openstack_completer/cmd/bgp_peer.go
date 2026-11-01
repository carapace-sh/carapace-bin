package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_peerCmd = &cobra.Command{
	Use:   "peer",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_peerCmd).Standalone()

	bgpCmd.AddCommand(bgp_peerCmd)
}
