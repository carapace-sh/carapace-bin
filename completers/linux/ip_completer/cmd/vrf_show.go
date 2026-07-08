package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vrf_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "show all configured VRF",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vrf_showCmd).Standalone()

	vrfCmd.AddCommand(vrf_showCmd)
}
