package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing connection profile in an interactive editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_editCmd).Standalone()

	connectionCmd.AddCommand(connection_editCmd)

	carapace.Gen(connection_editCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("id", "uuid", "path", "type", "con-name"),
			net.ActionConnections(),
		).ToA(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[0] {
			case "id", "uuid", "path":
				return net.ActionConnections()
			case "type":
				return ActionTypes()
			default:
				return carapace.ActionValues()
			}
		}),
	)

	carapace.Gen(connection_editCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "id", "uuid", "path":
					return net.ActionConnections()
				case "type":
					return ActionTypes()
				}
			}
			if len(c.Args)%2 == 0 {
				return carapace.ActionValues("id", "uuid", "path", "type", "con-name")
			}
			return carapace.ActionValues()
		}),
	)
}
