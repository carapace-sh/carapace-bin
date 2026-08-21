package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consistency_group_remove_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Remove volume(s) from consistency group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_remove_volumeCmd).Standalone()

	consistency_group_removeCmd.AddCommand(consistency_group_remove_volumeCmd)
}
