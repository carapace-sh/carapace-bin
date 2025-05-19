package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var patchIdCmd = &cobra.Command{
	Use:     "patch-id",
	Short:   "Compute unique ID for a patch",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(patchIdCmd).Standalone()

	patchIdCmd.Flags().Bool("stable", false, "use the stable patch-id algorithm")
	patchIdCmd.Flags().Bool("unstable", false, "use the unstable patch-id algorithm")
	patchIdCmd.Flags().Bool("verbatim", false, "don't strip whitespace from the patch")
	rootCmd.AddCommand(patchIdCmd)
}
