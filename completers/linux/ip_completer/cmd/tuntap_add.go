package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tuntap_addCmd = &cobra.Command{
	Use:   "add",
	Short: "create a TUN or TAP device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tuntap_addCmd).Standalone()

	tuntapCmd.AddCommand(tuntap_addCmd)

	carapace.Gen(tuntap_addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "mode":
					return carapace.ActionValues("tap", "tun")
				case "dev":
					return carapace.ActionValues()
				case "user":
					return carapace.ActionValues()
				case "group":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("dev", "group", "mode", "multi_queue", "one_queue", "pi", "user", "vnet_hdr")
		}),
	)
}
