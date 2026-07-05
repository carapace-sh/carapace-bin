package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_downCmd = &cobra.Command{
	Use:   "down",
	Short: "Deactivate a connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_downCmd).Standalone()

	connectionCmd.AddCommand(connection_downCmd)

	carapace.Gen(connection_downCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("id", "uuid", "path", "apath"),
			net.ActionConnections(),
		).ToA(),
	)

	carapace.Gen(connection_downCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "id", "uuid", "path", "apath":
					return net.ActionConnections()
				}
			}
			return carapace.Batch(
				carapace.ActionValues("id", "uuid", "path", "apath"),
				net.ActionConnections(),
			).ToA()
		}),
	)
}
