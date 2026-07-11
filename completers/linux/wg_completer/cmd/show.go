package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows the current configuration and device information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()
	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).PositionalCompletion(
		ActionInterfaces(),
		carapace.ActionValues(
			"public-key",
			"private-key",
			"listen-port",
			"fwmark",
			"peers",
			"preshared-keys",
			"endpoints",
			"allowed-ips",
			"latest-handshakes",
			"persistent-keepalive",
			"transfer",
			"dump",
		),
	)
}