package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_gradleCmd = &cobra.Command{
	Use:   "gradle <gradle args>",
	Short: "Run Gradle through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_gradleCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_gradleCmd)

	carapace.Gen(dependencyFirewall_gradleCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("gradle").Split(),
	)

}
