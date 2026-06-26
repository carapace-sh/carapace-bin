package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var discussion_commentCmd = &cobra.Command{
	Use:     "comment {<number> | <discussion-url> | <comment-id> | <comment-url>} [flags]",
	Short:   "Add, edit, or delete a comment or a reply on a discussion (preview)",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discussion_commentCmd).Standalone()

	discussion_commentCmd.Flags().StringP("body", "b", "", "Comment body text")
	discussion_commentCmd.Flags().StringP("body-file", "F", "", "Read body text from file (use \"-\" to read from standard input)")
	discussion_commentCmd.Flags().Bool("delete", false, "Delete the specified comment")
	discussion_commentCmd.Flags().Bool("edit", false, "Edit the specified comment")
	discussion_commentCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	discussion_commentCmd.Flags().Bool("yes", false, "Skip the delete confirmation prompt")
	discussionCmd.AddCommand(discussion_commentCmd)

	carapace.Gen(discussion_commentCmd).FlagCompletion(carapace.ActionMap{
		"body-file": carapace.ActionFiles(),
		"repo":      gh.ActionHostOwnerRepositories(),
	})

	carapace.Gen(discussion_commentCmd).PositionalAnyCompletion(
		action.ActionDiscussions(discussion_commentCmd),
	)
}
