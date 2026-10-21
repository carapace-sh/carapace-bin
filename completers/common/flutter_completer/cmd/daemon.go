package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run a persistent, JSON-RPC based server to communicate with devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemonCmd).Standalone()

	daemonCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	daemonCmd.Flags().String("listen-on-tcp-port", "", "If specified, the daemon will be listening for commands on the specified port instead of stdio.")
	rootCmd.AddCommand(daemonCmd)
}
