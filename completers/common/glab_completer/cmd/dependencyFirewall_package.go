package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependencyFirewall_packageCmd = &cobra.Command{
	Use:   "package <purl>",
	Short: "Check a package against the GitLab Dependency Firewall. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_packageCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_packageCmd)
}
