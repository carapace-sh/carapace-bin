package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_tier_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve an app tier",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_tier_getCmd).Standalone()
	apps_tierCmd.AddCommand(apps_tier_getCmd)
}
