package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consistency_group_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_removeCmd).Standalone()

	consistency_groupCmd.AddCommand(consistency_group_removeCmd)
}
