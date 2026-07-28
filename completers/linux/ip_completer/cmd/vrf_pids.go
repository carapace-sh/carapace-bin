package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vrf_pidsCmd = &cobra.Command{
	Use:   "pids",
	Short: "report processes associated with the named VRF",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vrf_pidsCmd).Standalone()

	vrfCmd.AddCommand(vrf_pidsCmd)
}
