package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_l3_conntrackCmd = &cobra.Command{
	Use:   "conntrack",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_l3_conntrackCmd).Standalone()

	network_l3Cmd.AddCommand(network_l3_conntrackCmd)
}
