package cmd

import (
	"github.com/spf13/cobra"
)

var patch_idCmd = &cobra.Command{
	Use:     "patch-id",
	Short:   "Compute unique ID for a patch",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {

	rootCmd.AddCommand(patch_idCmd)
}
