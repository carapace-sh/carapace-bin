package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_pipenvCmd = &cobra.Command{
	Use:   "pipenv <pipenv args>",
	Short: "Run Pipenv through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_pipenvCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_pipenvCmd)

	carapace.Gen(dependencyFirewall_pipenvCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("pipenv").Split(),
	)

}
