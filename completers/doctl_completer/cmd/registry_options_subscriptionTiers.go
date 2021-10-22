package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_options_subscriptionTiersCmd = &cobra.Command{
	Use:   "subscription-tiers",
	Short: "List available container registry subscription tiers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_options_subscriptionTiersCmd).Standalone()
	registry_optionsCmd.AddCommand(registry_options_subscriptionTiersCmd)
}
