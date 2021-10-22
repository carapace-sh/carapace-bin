package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_tier_instanceSize_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve an app instance size",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_tier_instanceSize_getCmd).Standalone()
	apps_tier_instanceSizeCmd.AddCommand(apps_tier_instanceSize_getCmd)
}
