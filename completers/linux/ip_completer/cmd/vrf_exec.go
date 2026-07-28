package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vrf_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "run a command against the named VRF",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vrf_execCmd).Standalone()

	vrfCmd.AddCommand(vrf_execCmd)
}
