package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var patch_idCmd = &cobra.Command{
	Use:     "patch-id",
	Short:   "Compute unique ID for a patch",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(patch_idCmd).Standalone()

	rootCmd.AddCommand(patch_idCmd)
}
