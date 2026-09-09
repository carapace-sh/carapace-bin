package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var region_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set region properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(region_setCmd).Standalone()

	region_setCmd.Flags().String("description", "", "New region description")
	region_setCmd.Flags().String("parent-region", "", "New parent region ID")
	regionCmd.AddCommand(region_setCmd)
}
