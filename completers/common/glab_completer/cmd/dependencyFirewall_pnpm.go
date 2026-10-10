package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_pnpmCmd = &cobra.Command{
	Use:   "pnpm <pnpm args>",
	Short: "Run pnpm through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_pnpmCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_pnpmCmd)

	carapace.Gen(dependencyFirewall_pnpmCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("pnpm").Split(),
	)

}
