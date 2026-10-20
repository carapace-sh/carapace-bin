package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap___applyFirewallPlanCmd = &cobra.Command{
	Use:    "__apply-firewall-plan",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap___applyFirewallPlanCmd).Standalone()

	bootstrap___applyFirewallPlanCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap___applyFirewallPlanCmd)
}
