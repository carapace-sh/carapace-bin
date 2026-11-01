package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_npmCmd = &cobra.Command{
	Use:   "npm <npm args>",
	Short: "Run npm through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_npmCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_npmCmd)

	carapace.Gen(dependencyFirewall_npmCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("npm").Split(),
	)

}
