package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var portForwardCmd = &cobra.Command{
	Use:     "port-forward TYPE/NAME [options] [LOCAL_PORT:]REMOTE_PORT [...[LOCAL_PORT_N:]REMOTE_PORT_N]",
	Short:   "Forward one or more local ports to a pod",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(portForwardCmd).Standalone()

	portForwardCmd.Flags().StringSlice("address", nil, "Addresses to listen on (comma separated). Only accepts IP addresses or localhost as a value. When localhost is supplied, kubectl will try to bind on both 127.0.0.1 and ::1 and will fail if neither of these addresses are available to bind.")
	portForwardCmd.Flags().String("pod-running-timeout", "", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	rootCmd.AddCommand(portForwardCmd)

	carapace.Gen(portForwardCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionApiResourceResources(kubectl.ApiResourceResourcesOpts{
				Namespace: rootCmd.Flag("namespace").Value.String(),
			})
		}),
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionPorts().Invoke(c).Suffix(":").ToA()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
