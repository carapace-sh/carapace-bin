package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_pipCmd = &cobra.Command{
	Use:   "pip <pip args>",
	Short: "Run Pip through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_pipCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_pipCmd)

	carapace.Gen(dependencyFirewall_pipCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("pip").Split(),
	)

}
