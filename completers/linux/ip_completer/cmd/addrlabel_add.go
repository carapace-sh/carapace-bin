package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var addrlabel_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add an address label entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addrlabel_addCmd).Standalone()

	addrlabelCmd.AddCommand(addrlabel_addCmd)

	carapace.Gen(addrlabel_addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "prefix":
					return net.ActionSubnets()
				case "dev":
					return net.ActionDevices(net.IncludedDevices{Wifi: true, Ethernet: true})
				case "label":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("prefix", "dev", "label")
		}),
	)
}
