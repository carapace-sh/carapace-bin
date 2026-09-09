package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consistency_group_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set consistency group properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_setCmd).Standalone()

	consistency_group_setCmd.Flags().String("description", "", "New consistency group description")
	consistency_group_setCmd.Flags().String("name", "", "New consistency group name")
	consistency_groupCmd.AddCommand(consistency_group_setCmd)
}
