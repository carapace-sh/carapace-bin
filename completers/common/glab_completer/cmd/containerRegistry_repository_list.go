package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var containerRegistry_repository_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List container registry repositories.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerRegistry_repository_listCmd).Standalone()

	containerRegistry_repository_listCmd.PersistentFlags().StringP("group", "g", "", "List container registry repositories for a group.")
	containerRegistry_repository_listCmd.Flags().Bool("include-tag-details", false, "Fetch digest, size, and creation time for included tags. Makes one API call per tag. Project JSON output only. Implies --include-tags.")
	containerRegistry_repository_listCmd.Flags().Bool("include-tags", false, "Include tags in the response. Project repositories only.")
	containerRegistry_repository_listCmd.Flags().Bool("include-tags-count", false, "Include the number of tags in the response. Project repositories only.")
	containerRegistry_repository_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	containerRegistry_repository_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	containerRegistry_repository_listCmd.Flags().StringP("page", "p", "", "Page number.")
	containerRegistry_repository_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	containerRegistry_repository_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	containerRegistry_repositoryCmd.AddCommand(containerRegistry_repository_listCmd)

	carapace.Gen(containerRegistry_repository_listCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(containerRegistry_repository_listCmd),
	})
}
