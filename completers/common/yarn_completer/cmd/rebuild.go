package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rebuildCmd = &cobra.Command{
	Use:     "rebuild",
	Short:   "Rebuild the project's native packages",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebuildCmd).Standalone()

	rootCmd.AddCommand(rebuildCmd)

	// TODO positional completion
}
