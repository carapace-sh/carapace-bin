package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a private container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_createCmd).Standalone()
	registry_createCmd.Flags().String("subscription-tier", "basic", "Subscription tier for the new registry. Possible values: see `doctl registry options subscription-tiers` (required)")
	registryCmd.AddCommand(registry_createCmd)
}
