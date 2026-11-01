package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor an existing connection profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_monitorCmd).Standalone()

	connectionCmd.AddCommand(connection_monitorCmd)

	carapace.Gen(connection_monitorCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("id", "path", "uuid"),
			net.ActionConnections(),
		).ToA(),
	)

	carapace.Gen(connection_monitorCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "id", "uuid", "path":
					return net.ActionConnections()
				}
			}
			return carapace.Batch(
				carapace.ActionValues("id", "path", "uuid"),
				net.ActionConnections(),
			).ToA()
		}),
	)
}
