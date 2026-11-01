package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_poetryCmd = &cobra.Command{
	Use:   "poetry <poetry args>",
	Short: "Run Poetry through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_poetryCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_poetryCmd)

	carapace.Gen(dependencyFirewall_poetryCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("poetry").Split(),
	)

}
