package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_upCmd = &cobra.Command{
	Use:   "up",
	Short: "Activate a connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_upCmd).Standalone()

	connectionCmd.AddCommand(connection_upCmd)

	carapace.Gen(connection_upCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("id", "uuid", "path"),
			net.ActionConnections(),
		).ToA(),
	)

	carapace.Gen(connection_upCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "id", "uuid", "path":
					return net.ActionConnections()
				case "ifname":
					return net.ActionDevices(net.AllDevices)
				case "ap":
					return net.ActionBssids()
				case "passwd-file":
					return carapace.ActionFiles()
				}
			}
			return carapace.ActionValues("ifname", "ap", "passwd-file")
		}),
	)
}
