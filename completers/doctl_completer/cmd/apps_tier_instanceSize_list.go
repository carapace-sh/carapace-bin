package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_tier_instanceSize_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all app instance sizes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_tier_instanceSize_listCmd).Standalone()
	apps_tier_instanceSizeCmd.AddCommand(apps_tier_instanceSize_listCmd)
}
