package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var discussion_editCmd = &cobra.Command{
	Use:     "edit {<number> | <discussion-url>} [flags]",
	Short:   "Edit a discussion (preview)",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discussion_editCmd).Standalone()

	discussion_editCmd.Flags().StringSlice("add-label", nil, "Add labels by `name`")
	discussion_editCmd.Flags().StringP("body", "b", "", "New body for the discussion")
	discussion_editCmd.Flags().StringP("body-file", "F", "", "Read body text from file (use \"-\" to read from standard input)")
	discussion_editCmd.Flags().StringP("category", "c", "", "New category name or slug for the discussion")
	discussion_editCmd.Flags().StringSlice("remove-label", nil, "Remove labels by `name`")
	discussion_editCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	discussion_editCmd.Flags().StringP("title", "t", "", "New title for the discussion")
	discussionCmd.AddCommand(discussion_editCmd)

	carapace.Gen(discussion_editCmd).FlagCompletion(carapace.ActionMap{
		"add-label":    action.ActionLabels(discussion_editCmd).UniqueList(","),
		"body-file":    carapace.ActionFiles(),
		"category":     action.ActionDiscussionCategories(discussion_editCmd),
		"remove-label": action.ActionLabels(discussion_editCmd).UniqueList(","),
		"repo":         gh.ActionHostOwnerRepositories(),
	})

	carapace.Gen(discussion_editCmd).PositionalAnyCompletion(
		action.ActionDiscussions(discussion_editCmd),
	)
}
