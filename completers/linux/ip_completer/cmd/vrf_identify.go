package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vrf_identifyCmd = &cobra.Command{
	Use:   "identify",
	Short: "report VRF association for a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vrf_identifyCmd).Standalone()

	vrfCmd.AddCommand(vrf_identifyCmd)
}
