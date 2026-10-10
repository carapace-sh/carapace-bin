package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_gemCmd = &cobra.Command{
	Use:   "gem <gem args>",
	Short: "Run gem through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_gemCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_gemCmd)

	carapace.Gen(dependencyFirewall_gemCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("gem").Split(),
	)

}
