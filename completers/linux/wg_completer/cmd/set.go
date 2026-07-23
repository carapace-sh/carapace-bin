package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/wg_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Change the current configuration, add peers, remove peers, or change peers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).PositionalCompletion(
		action.ActionInterfaces(),
	)

	carapace.Gen(setCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			previous := c.Args[len(c.Args)-1]
			switch previous {
			case "listen-port":
				return net.ActionPorts()
			case "fwmark":
				return carapace.ActionValues("off")
			case "private-key", "preshared-key":
				return carapace.ActionFiles()
			case "peer":
				return carapace.ActionValues()
			case "endpoint":
				return carapace.ActionValues()
			case "persistent-keepalive":
				return carapace.ActionValues("off")
			case "allowed-ips":
				return carapace.ActionValues()
			case "remove":
				return carapace.ActionValues(
					"fwmark",
					"listen-port",
					"peer",
					"private-key",
				)
			default:
				if isPeerKey(c.Args[1:]) {
					return carapace.ActionValues(
						"allowed-ips",
						"endpoint",
						"persistent-keepalive",
						"preshared-key",
						"remove",
					)
				}
				return carapace.ActionValues(
					"fwmark",
					"listen-port",
					"peer",
					"private-key",
				)
			}
		}),
	)
}

func isPeerKey(args []string) bool {
	for i := len(args) - 1; i >= 0; i-- {
		switch args[i] {
		case "peer":
			return true
		case "listen-port", "fwmark", "private-key":
			return false
		}
	}
	return false
}
