package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gc_repackCmd = &cobra.Command{
	Use:   "repack",
	Short: "Repack reachable objects into one pack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gc_repackCmd).Standalone()

	gc_repackCmd.Flags().Bool("bitmaps", false, "Write a reachability `.bitmap` sidecar for the new pack")
	gc_repackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gcCmd.AddCommand(gc_repackCmd)
}
