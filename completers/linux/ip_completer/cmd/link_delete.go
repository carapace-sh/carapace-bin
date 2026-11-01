package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var link_deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete a virtual link",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_deleteCmd).Standalone()

	linkCmd.AddCommand(link_deleteCmd)

	carapace.Gen(link_deleteCmd).PositionalCompletion(
		net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true}),
	)
}
