package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_tier_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all app tiers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_tier_listCmd).Standalone()
	apps_tierCmd.AddCommand(apps_tier_listCmd)
}
