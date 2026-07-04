package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ipconfig",
	Short: "display all current TCP/IP network configuration values",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/ipconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"all", "display full configuration information",
			"release", "release the IP address",
			"renew", "renew the IP address",
			"flushdns", "purge the DNS resolver cache",
			"registerdns", "initiate manual dynamic registration",
			"displaydns", "display DNS resolver cache contents",
			"showclassid", "display DHCP class ID for adapter",
			"setclassid", "set DHCP class ID for adapter",
		),
	)
}
