package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_dangerouslyDeleteCmd = &cobra.Command{
	Use:   "dangerously-delete",
	Short: "Dangerously delete all cached content by tag",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_dangerouslyDeleteCmd).Standalone()

	cache_dangerouslyDeleteCmd.Flags().String("project", "", "Project name or ID")
	cache_dangerouslyDeleteCmd.Flags().String("revalidation-deadline-seconds", "", "Revalidation deadline in seconds")
	cache_dangerouslyDeleteCmd.Flags().String("srcimg", "", "Source image URL")
	cache_dangerouslyDeleteCmd.Flags().String("tag", "", "Cache tag to delete")
	cache_dangerouslyDeleteCmd.Flags().Bool("yes", false, "Skip confirmation")

	cacheCmd.AddCommand(cache_dangerouslyDeleteCmd)
}
