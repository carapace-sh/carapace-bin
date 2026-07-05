package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate connection profiles to a different settings plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_migrateCmd).Standalone()

	connection_migrateCmd.Flags().StringSlice("plugin", nil, "settings plugin to migrate to")
	connectionCmd.AddCommand(connection_migrateCmd)

	carapace.Gen(connection_migrateCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("id", "uuid", "path"),
			net.ActionConnections(),
		).ToA(),
	)

	carapace.Gen(connection_migrateCmd).PositionalAnyCompletion(
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
