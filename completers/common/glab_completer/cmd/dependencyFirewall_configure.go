package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependencyFirewall_configureCmd = &cobra.Command{
	Use:   "configure <package-manager>",
	Short: "Configure Dependency Firewall registry URLs for a package manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyFirewall_configureCmd).Standalone()

	dependencyFirewall_configureCmd.Flags().String("repo-deploy", "", "Full registry URL to deploy (publish) packages to.")
	dependencyFirewall_configureCmd.Flags().String("repo-resolve", "", "Full registry URL to resolve (install) packages from.")
	dependencyFirewallCmd.AddCommand(dependencyFirewall_configureCmd)
}
