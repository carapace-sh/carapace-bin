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
				case "dev":
					return carapace.ActionValues()
				case "dsfield":
					return carapace.ActionValues()
				case "encaplimit":
					return carapace.ActionValues()
				case "flowlabel":
					return carapace.ActionValues()
				case "hoplimit":
					return carapace.ActionValues()
				case "ikey":
					return carapace.ActionValues()
				case "key":
					return carapace.ActionValues()
				case "local":
					return carapace.ActionValues().NoSpace()
				case "mode":
					return carapace.ActionValues("any", "gre", "ip6gre", "ip6ip6", "ipip", "ipip6", "isatap", "sit", "vti", "vti6")
				case "okey":
					return carapace.ActionValues()
				case "remote":
					return carapace.ActionValues().NoSpace()
				case "tclass":
					return carapace.ActionValues()
				case "tos":
					return carapace.ActionValues()
				case "ttl":
					return carapace.ActionValues()
				}
			}
			return carapace.ActionValues("allow-localremote", "csum", "dev", "dsfield", "encaplimit", "flowlabel", "hoplimit", "icsum", "ignore-df", "ikey", "key", "local", "mode", "nopmtudisc", "ocsum", "okey", "remote", "tclass", "tos", "ttl")
		}),
	)
}
