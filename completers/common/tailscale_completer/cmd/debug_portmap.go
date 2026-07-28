package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_portmapCmd = &cobra.Command{
	Use:   "portmap",
	Short: "Run portmap debugging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_portmapCmd).Standalone()

	debug_portmapCmd.Flags().String("duration", "5s", "timeout for port mapping")
	debug_portmapCmd.Flags().String("gateway-addr", "", "override gateway IP (must also pass --self-addr)")
	debug_portmapCmd.Flags().Bool("log-http", false, "print all HTTP requests and responses to the log")
	debug_portmapCmd.Flags().String("self-addr", "", "override self IP (must also pass --gateway-addr)")
	debug_portmapCmd.Flags().String("type", "", "portmap debug type (one of \"\", \"pmp\", \"pcp\", or \"upnp\")")
	debugCmd.AddCommand(debug_portmapCmd)

	carapace.Gen(debug_portmapCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("pmp", "pcp", "upnp"),
	})
}
