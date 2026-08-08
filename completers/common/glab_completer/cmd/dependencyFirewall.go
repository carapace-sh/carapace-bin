package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependencyFirewallCmd = &cobra.Command{
	Use:     "dependency-firewall <command>",
	Short:   "Configure and monitor GitLab Dependency Firewall for local package managers.",
	Aliases: []string{"df"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewallCmd).Standalone()

	rootCmd.AddCommand(dependencyFirewallCmd)
}
