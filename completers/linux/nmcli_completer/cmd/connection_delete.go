package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing connection profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_deleteCmd).Standalone()

	connectionCmd.AddCommand(connection_deleteCmd)

	carapace.Gen(connection_deleteCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("id", "uuid", "path"),
			net.ActionConnections(),
		).ToA(),
	)

	carapace.Gen(connection_deleteCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "id", "uuid", "path":
					return net.ActionConnections()
				}
			}
			return carapace.Batch(
				carapace.ActionValues("id", "uuid", "path"),
				net.ActionConnections(),
			).ToA()
		}),
	)
}
