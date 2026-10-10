package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dependencyFirewall_uvCmd = &cobra.Command{
	Use:   "uv <uv args>",
	Short: "Run uv through the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_uvCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_uvCmd)

	carapace.Gen(dependencyFirewall_uvCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("uv").Split(),
	)

}
