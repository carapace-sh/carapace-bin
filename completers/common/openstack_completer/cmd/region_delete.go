package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var region_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete region(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(region_deleteCmd).Standalone()

	regionCmd.AddCommand(region_deleteCmd)
}
