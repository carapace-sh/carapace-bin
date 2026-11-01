package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_twineCmd = &cobra.Command{
	Use:   "twine <twine args>",
	Short: "Run Twine through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_twineCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_twineCmd)

	carapace.Gen(dependencyFirewall_twineCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("twine").Split(),
	)

}
