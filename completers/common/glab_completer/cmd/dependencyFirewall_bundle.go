package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_bundleCmd = &cobra.Command{
	Use:   "bundle <bundle args>",
	Short: "Run Bundler through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_bundleCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_bundleCmd)

	carapace.Gen(dependencyFirewall_bundleCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("bundle").Split(),
	)

}
