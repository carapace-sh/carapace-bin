package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_repository_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "List tags for a repository in a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_repository_listTagsCmd).Standalone()
	registry_repository_listTagsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Tag`, `CompressedSizeBytes`, `UpdatedAt`, `ManifestDigest`")
	registry_repository_listTagsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	registry_repositoryCmd.AddCommand(registry_repository_listTagsCmd)
}
