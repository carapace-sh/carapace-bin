package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dryBuildCmd = &cobra.Command{
	Use:   "dry-build",
	Short: "Show what store paths would be built or downloaded, but otherwise do nothing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dryBuildCmd).Standalone()
	rootCmd.AddCommand(dryBuildCmd)
}
