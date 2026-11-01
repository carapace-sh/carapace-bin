package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var CleanupMountpointsCmd = &cobra.Command{
	Use:   "Cleanup-Mountpoints",
	Short: "delete resources associated with corrupted mounted images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(CleanupMountpointsCmd).Standalone()
	rootCmd.AddCommand(CleanupMountpointsCmd)
}
