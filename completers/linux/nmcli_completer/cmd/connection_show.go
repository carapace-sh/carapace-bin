package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var connection_showCmd = &cobra.Command{
	Use:   "show",
	Short: "show details for specified connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_showCmd).Standalone()

	connection_showCmd.Flags().Bool("active", false, "show active connections")
	connection_showCmd.Flags().String("order", "", "specify order")

	connectionCmd.AddCommand(connection_showCmd)

	carapace.Gen(connection_showCmd).FlagCompletion(carapace.ActionMap{
		"order": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("+active", "-active", "+name", "-name", "+type", "-type", "+path", "-path", "active", "name", "type", "path").Suffix(":")
			default:
				return carapace.ActionValues("active", "name", "type", "path")
			}
		}),
	})

	carapace.Gen(connection_showCmd).PositionalAnyCompletion(
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
