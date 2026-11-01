package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_yarnCmd = &cobra.Command{
	Use:   "yarn <yarn args>",
	Short: "Run Yarn through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_yarnCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_yarnCmd)

	carapace.Gen(dependencyFirewall_yarnCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("yarn").Split(),
	)

}
