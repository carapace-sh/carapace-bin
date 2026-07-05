package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var device_lldpCmd = &cobra.Command{
	Use:   "lldp",
	Short: "Display information about neighboring devices learned through LLDP",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_lldpCmd).Standalone()

	deviceCmd.AddCommand(device_lldpCmd)

	carapace.Gen(device_lldpCmd).PositionalCompletion(
		carapace.ActionValues("list", "ifname"),
	)

	carapace.Gen(device_lldpCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "ifname":
					return net.ActionDevices(net.AllDevices)
				}
			}
			return carapace.ActionValues("list", "ifname")
		}),
	)
}
