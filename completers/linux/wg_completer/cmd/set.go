package cmd

import (
	"github.com/carapace-sh/carapace"
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
	carapace.Gen(setCmd).PositionalCompletion(ActionInterfaces())
	carapace.Gen(setCmd).PositionalAnyCompletion(
		carapace.ActionValues(
			"listen-port",
			"fwmark",
			"private-key",
			"peer",
			"remove",
			"preshared-key",
			"endpoint",
			"persistent-keepalive",
			"allowed-ips",
		),
	)
}