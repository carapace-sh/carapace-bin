package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_volume_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a block storage volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_volume_deleteCmd).Standalone()
	compute_volume_deleteCmd.Flags().BoolP("force", "f", false, "Force volume delete")
	compute_volumeCmd.AddCommand(compute_volume_deleteCmd)
}
