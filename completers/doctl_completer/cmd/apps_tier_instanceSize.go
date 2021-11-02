package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_tier_instanceSizeCmd = &cobra.Command{
	Use:   "instance-size",
	Short: "Display commands for working with app instance sizes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_tier_instanceSizeCmd).Standalone()
	apps_tierCmd.AddCommand(apps_tier_instanceSizeCmd)
}
