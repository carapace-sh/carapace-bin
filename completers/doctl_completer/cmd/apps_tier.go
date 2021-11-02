package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_tierCmd = &cobra.Command{
	Use:   "tier",
	Short: "Display commands for working with app tiers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_tierCmd).Standalone()
	appsCmd.AddCommand(apps_tierCmd)
}
