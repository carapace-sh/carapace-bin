package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_mavenCmd = &cobra.Command{
	Use:   "maven <mvn args>",
	Short: "Run Maven through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_mavenCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_mavenCmd)

	carapace.Gen(dependencyFirewall_mavenCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("mvn").Split(),
	)

}
