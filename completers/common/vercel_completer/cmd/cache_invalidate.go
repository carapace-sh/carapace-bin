package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_invalidateCmd = &cobra.Command{
	Use:   "invalidate",
	Short: "Invalidate all cached content by tag",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_invalidateCmd).Standalone()

	cache_invalidateCmd.Flags().String("project", "", "Project name or ID")
	cache_invalidateCmd.Flags().String("srcimg", "", "Source image URL")
	cache_invalidateCmd.Flags().String("tag", "", "Cache tag to invalidate")
	cache_invalidateCmd.Flags().Bool("yes", false, "Skip confirmation")

	cacheCmd.AddCommand(cache_invalidateCmd)
}
