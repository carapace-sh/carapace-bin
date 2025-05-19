package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var codespace_ports_forwardCmd = &cobra.Command{
	Use:   "forward <remote-port>:<local-port>...",
	Short: "Forward ports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_ports_forwardCmd).Standalone()

	codespace_portsCmd.AddCommand(codespace_ports_forwardCmd)

	carapace.Gen(codespace_ports_forwardCmd).PositionalCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				codespace := codespace_ports_forwardCmd.Flag("codespace").Value.String()
				return action.ActionCodespacePorts(codespace).Invoke(c).Suffix(":").ToA()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
