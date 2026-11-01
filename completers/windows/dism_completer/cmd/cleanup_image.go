package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var CleanupImageCmd = &cobra.Command{
	Use:   "Cleanup-Image",
	Short: "perform cleanup and recovery on an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(CleanupImageCmd).Standalone()
	rootCmd.AddCommand(CleanupImageCmd)
}
