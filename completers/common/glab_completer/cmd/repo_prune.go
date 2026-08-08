package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_pruneCmd = &cobra.Command{
	Use:   "prune [flags]",
	Short: "Delete local Git branches whose merge request has been merged.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_pruneCmd).Standalone()

	repo_pruneCmd.Flags().Bool("dry-run", false, "Preview branches that would be deleted without deleting them.")
	repo_pruneCmd.Flags().StringSliceP("exclude", "e", nil, "Branch name or glob pattern to exclude. Comma-separated or repeated.")
	repo_pruneCmd.Flags().Bool("merged", false, "Use 'git branch --merged' instead of querying GitLab. Detects fast-forward merges only.")
	repo_pruneCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt.")
	repoCmd.AddCommand(repo_pruneCmd)

	carapace.Gen(repo_pruneCmd).FlagCompletion(carapace.ActionMap{
		"exclude": action.ActionBranches(repo_pruneCmd),
	})
}
