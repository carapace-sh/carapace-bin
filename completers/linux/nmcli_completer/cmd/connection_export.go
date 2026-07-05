package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_exportCmd).Standalone()

	connectionCmd.AddCommand(connection_exportCmd)

	carapace.Gen(connection_exportCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("id", "uuid", "path"),
			net.ActionConnections(),
		).ToA(),
		net.ActionConnections(),
		carapace.ActionFiles(),
	)

	carapace.Gen(connection_exportCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "id", "uuid", "path":
					return net.ActionConnections()
				}
			}
			if len(c.Args)%2 == 0 && len(c.Args) >= 2 {
				return carapace.ActionFiles()
			}
			return carapace.ActionValues("id", "uuid", "path")
		}),
	)
}
