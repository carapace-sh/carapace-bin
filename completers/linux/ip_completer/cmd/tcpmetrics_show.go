package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var tcpmetrics_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "show cached TCP metrics entries",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tcpmetrics_showCmd).Standalone()

	tcpmetricsCmd.AddCommand(tcpmetrics_showCmd)

	carapace.Gen(tcpmetrics_showCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "address":
					return net.ActionIpv4Addresses()
				case "dev":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				}
			}
			return carapace.ActionValues("address", "dev")
		}),
	)
}
