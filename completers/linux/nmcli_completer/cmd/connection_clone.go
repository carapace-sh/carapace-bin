package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone an existing connection profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_cloneCmd).Standalone()

	connection_cloneCmd.Flags().Bool("temporary", false, "only temporary")
	connectionCmd.AddCommand(connection_cloneCmd)

	carapace.Gen(connection_cloneCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("id", "uuid", "path"),
			net.ActionConnections(),
		).ToA(),
		net.ActionConnections(),
		carapace.ActionValues(), // new name
	)

	carapace.Gen(connection_cloneCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "id", "uuid", "path":
					return net.ActionConnections()
				}
			}
			if len(c.Args)%2 == 0 && len(c.Args) >= 2 {
				return carapace.ActionValues() // new name is freeform
			}
			return carapace.ActionValues("id", "uuid", "path")
		}),
	)
}
