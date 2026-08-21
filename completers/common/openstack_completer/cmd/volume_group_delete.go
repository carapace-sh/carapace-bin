package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a volume group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_deleteCmd).Standalone()

	volume_group_deleteCmd.Flags().Bool("force", false, "Delete the volume group even if it contains volumes.")
	volume_groupCmd.AddCommand(volume_group_deleteCmd)
}
