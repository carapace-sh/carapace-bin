package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tunnel_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new tunnel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tunnel_addCmd).Standalone()

	tunnelCmd.AddCommand(tunnel_addCmd)

	carapace.Gen(tunnel_addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[len(c.Args)-1] {
				case "mode":
					return carapace.ActionValues("ipip", "gre", "sit", "isatap", "vti", "ip6ip6", "ipip6", "ip6gre", "vti6", "any")
				case "remote":
					return carapace.ActionValues().NoSpace()
				case "local":
					return carapace.ActionValues().NoSpace()
				case "dev":
					return carapace.ActionValues()
				case "encaplimit":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("mode", "remote", "local", "dev", "ttl", "hoplimit", "tos", "dsfield", "tclass", "nopmtudisc", "ignore-df", "key", "ikey", "okey", "csum", "icsum", "ocsum", "encaplimit", "flowlabel", "allow-localremote")
		}),
	)
}
