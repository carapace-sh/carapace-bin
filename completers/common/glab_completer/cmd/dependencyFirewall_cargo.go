package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_cargoCmd = &cobra.Command{
	Use:   "cargo <cargo args>",
	Short: "Run Cargo through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_cargoCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_cargoCmd)

	carapace.Gen(dependencyFirewall_cargoCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("cargo").Split(),
	)
}
