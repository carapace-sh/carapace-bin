package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependencyFirewall_ciSummaryCmd = &cobra.Command{
	Use:   "ci-summary",
	Short: "Summarize Dependency Firewall activity from the CI log.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_ciSummaryCmd).Standalone()

	dependencyFirewallCmd.AddCommand(dependencyFirewall_ciSummaryCmd)
}
