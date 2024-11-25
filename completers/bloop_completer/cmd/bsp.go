package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bloop"
	"github.com/spf13/cobra"
)

var bspCmd = &cobra.Command{
	Use:   "bsp",
	Short: "start build server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bspCmd).Standalone()

	bspCmd.Flags().String("host", "", "the server host for the bsp server")
	bspCmd.Flags().String("pipe-name", "", "a path to a new existing socket file to communicate through")
	bspCmd.Flags().String("port", "", "the port for the bsp server")
	bspCmd.Flags().String("protocol", "", "the connection protocol for the bsp server")
	bspCmd.Flags().String("socket", "", "a path to a socket file to communicate through")
	rootCmd.AddCommand(bspCmd)

	carapace.Gen(bspCmd).FlagCompletion(carapace.ActionMap{
		"host":      net.ActionHosts(),
		"pipe-name": carapace.ActionFiles(),
		"port":      net.ActionPorts(),
		"protocol":  bloop.ActionProtocols(),
		"socket":    carapace.ActionFiles(),
	})
}
